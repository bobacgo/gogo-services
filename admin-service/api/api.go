package api

import (
	"github.com/gogoclouds/gogo-services/common-lib/_examples/api/v1/demo"
	"google.golang.org/grpc"
)

func Router(server *grpc.Server) {
	demo.RegisterDemoServiceServer(server, &demo.DemoHandler{})
}
