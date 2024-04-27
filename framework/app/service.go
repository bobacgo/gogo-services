package app

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gogoclouds/gogo-services/framework/pkg/uid"

	"github.com/gogoclouds/gogo-services/framework/app/registry"
	"github.com/gogoclouds/gogo-services/framework/pkg/network"
)

type App struct {
	opts Options

	wg     sync.WaitGroup
	exit   chan struct{}
	signal chan os.Signal

	mu       sync.Mutex
	instance *registry.ServiceInstance
}

func New(opts ...Option) *App {
	o := Options{
		appid:           uid.UUID(),
		sigs:            []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registryTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&o)
	}
	return &App{
		opts:   o,
		exit:   make(chan struct{}),
		signal: make(chan os.Signal, 1),
	}
}

// Run run server
// 1.注册服务
// 2.退出相关组件或服务
func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}
	a.mu.Lock()
	a.instance = instance
	a.mu.Unlock()

	opts := a.opts

	if opts.httpServer != nil {
		go RunMustHttpServer(a, opts.httpServer)
	}

	if opts.rpcServer != nil {
		go RunMustRpcServer(a, opts.rpcServer)
	}

	ctx := context.Background()

	for _, fn := range opts.beforeStart {
		if err = fn(ctx); err != nil {
			return err
		}
	}

	// 注册服务 TODO 等待服务启动成功
	if opts.registrar != nil {
		ctx, cancel := context.WithTimeout(context.Background(), opts.registryTimeout)
		defer cancel()
		if err := opts.registrar.Registry(ctx, instance); err != nil {
			slog.Error("register service error", "err", err)
			return err
		}
	}

	for _, fn := range opts.afterStart {
		if err = fn(ctx); err != nil {
			return err
		}
	}

	slog.Info("server started")
	// 阻塞,监听退出信号
	signal.Notify(a.signal, opts.sigs...)
	<-a.signal

	err = a.shutdown(ctx)

	slog.Info("service has exited")
	return err
}

// Stop 手动停止服务
func (a *App) Stop() error {
	a.signal <- syscall.SIGTERM
	return nil
}

// shutdown server
// 1.注销服务
// 2.退出 http、grpc服务
func (a *App) shutdown(ctx context.Context) (err error) {
	opts := a.opts

	for _, fn := range opts.beforeStop {
		err = fn(ctx)
	}

	a.mu.Lock()
	instance := a.instance
	a.mu.Unlock()
	if opts.registrar != nil && instance != nil {
		ctx, cancel := context.WithTimeout(context.Background(), opts.registryTimeout)
		defer cancel()
		if err := opts.registrar.Deregister(ctx, instance); err != nil {
			slog.Error("deregister service error", "err", err)
			return err
		}
	}

	close(a.exit) // 通知http、rpc服务退出信号

	// 1.等待 Http 服务结束退出
	// 2.等待 RPC 服务结束退出
	a.wg.Wait()

	for _, fn := range opts.afterStop {
		err = fn(ctx)
	}

	return
}

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	opts := a.opts

	endpoints := make([]string, 0)
	httpScheme, grpcScheme := false, false
	for _, e := range opts.endpoints {
		switch strings.ToLower(e.Scheme) {
		case "https", "http":
			httpScheme = true
		case "grpc":
			grpcScheme = true
		}
		endpoints = append(endpoints, e.String())
	}
	if !httpScheme {
		if rUrl, err := getRegistryUrl("http", opts.conf.Server.Http.Addr); err == nil {
			endpoints = append(endpoints, rUrl)
		} else {
			slog.Error("get http registry err", "err", err)
		}
	}
	if !grpcScheme {
		if rUrl, err := getRegistryUrl("grpc", opts.conf.Server.Rpc.Addr); err == nil {
			endpoints = append(endpoints, rUrl)
		} else {
			slog.Error("get grpc registry err", "err", err)
		}
	}
	return &registry.ServiceInstance{
		ID:        opts.appid,
		Name:      opts.conf.Name,
		Version:   opts.conf.Version,
		Metadata:  nil,
		Endpoints: endpoints,
	}, nil
}

func getRegistryUrl(scheme, addr string) (string, error) {
	ip, err := network.OutBoundIP()
	if err != nil {
		return "", err
	}
	_, ports, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	return scheme + "://" + net.JoinHostPort(ip, ports), nil
}
