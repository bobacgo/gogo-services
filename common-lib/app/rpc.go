package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func RunRpcServer(app *App, register func(server *grpc.Server, app *Options)) {
	app.wg.Add(1)
	defer app.wg.Done()

	cfg := app.opts.GetConf()

	lis, err := net.Listen("tcp", cfg.Server.Rpc.Addr)
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer()

	// 注册健康检查服务
	healthgrpc.RegisterHealthServer(s, health.NewServer())

	if register != nil {
		register(s, &app.opts)
	}

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Panic(err)
		}
	}()

	<-app.exit       // 阻塞,等待被关闭
	s.GracefulStop() // 优雅停止
}

// RPC Dial

var rpcClientMap = make(map[string]*grpc.ClientConn)

func RpcDial(serverName string) (*grpc.ClientConn, error) {
	if cc, ok := rpcClientMap[serverName]; ok {
		state := cc.GetState()
		if state == connectivity.Ready {
			return cc, nil
		}
	}

	// conn, err := grpc.Dial(serverName, grpc.WithInsecure())
	conn, err := grpc.Dial(serverName, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	rpcClientMap[serverName] = conn
	return conn, nil
}
