package api

import (
	"github.com/gogoclouds/gogo-services/common-lib/_examples/api/v1/demo"
	"google.golang.org/grpc"
	"net/http"
)

func Router(h http.Handler) {
	server := h.(*grpc.Server)
	demo.RegisterDemoServiceServer(server, &demo.DemoHandler{})
}
