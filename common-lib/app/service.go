package app

import (
	"context"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uid"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/registry"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/network"
)

type App struct {
	Opts options

	wg   sync.WaitGroup
	exit chan struct{}

	mu       sync.Mutex
	instance *registry.ServiceInstance
}

func New(opts ...Option) *App {
	o := options{
		appid:           uid.UUID(),
		sigs:            []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registryTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&o)
	}
	return &App{
		Opts: o,
		exit: make(chan struct{}),
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

	opts := a.Opts

	if opts.httpServer != nil {
		go RunHttpServer(a, opts.httpServer)
	}

	if opts.rpcServer != nil {
		go RunRpcServer(a, opts.rpcServer)
	}

	ctx := context.Background()

	for _, fn := range opts.beforeStart {
		if err = fn(ctx); err != nil {
			return err
		}
	}

	// 注册服务
	if opts.registrar != nil {
		ctx, cancel := context.WithTimeout(context.Background(), opts.registryTimeout)
		defer cancel()
		if err := opts.registrar.Registry(ctx, instance); err != nil {
			logger.Errorf("register service error: %v", err)
			return err
		}
	}

	for _, fn := range opts.afterStart {
		if err = fn(ctx); err != nil {
			return err
		}
	}

	// 阻塞,监听退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, opts.sigs...)
	<-c

	err = a.Stop(ctx)

	logger.Info("service has exited")
	return err
}

// Stop stop server
// 1.注销服务
// 2.退出 http、grpc服务
func (a *App) Stop(ctx context.Context) (err error) {
	opts := a.Opts

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
			logger.Errorf("deregister service error: %w", err)
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
	opts := a.Opts

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
		if rUrl, err := getRegistryUrl("http", opts.Conf.Server.Http.Addr); err == nil {
			endpoints = append(endpoints, rUrl)
		} else {
			logger.Errorf("get http registry err:%v", err)
		}
	}
	if !grpcScheme {
		if rUrl, err := getRegistryUrl("grpc", opts.Conf.Server.Rpc.Addr); err == nil {
			endpoints = append(endpoints, rUrl)
		} else {
			logger.Errorf("get grpc registry err:%v", err)
		}
	}
	return &registry.ServiceInstance{
		ID:        opts.appid,
		Name:      opts.Conf.Name,
		Version:   opts.Conf.Version,
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
