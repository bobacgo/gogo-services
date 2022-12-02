package server

import (
	"context"
	"net/http"
	"testing"

	"github.com/gogoclouds/gogo-services/common-lib/_examples/api/v1/demo"

	"google.golang.org/grpc"
)

// pleaes run debug mode
func Test_RPCServer(t *testing.T) {
	RunRpcServer("0.0.0.0:9080", handle)
}

// pleaes run test mode
func Test_RPCDialPing(t *testing.T) {
	cc, err := RpcDial("localhost:9080")
	if err != nil {
		t.Fatal(err)
	}
	defer cc.Close()

	dsc := demo.NewDemoServiceClient(cc)
	res, err := dsc.Ping(context.Background(), &demo.PingRequest{
		Addr: "localhost.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 0 {
		t.Fatalf("code: %v, msg: %s", res.Code, res.Msg)
	}
	t.Log(res.Data)
}

// pleaes run test mode
func Test_RPCDialHi(t *testing.T) {
	cc, err := RpcDial("localhost:9080")
	if err != nil {
		t.Fatal(err)
	}
	defer cc.Close()

	dsc := demo.NewDemoServiceClient(cc)
	res, err := dsc.Hi(context.Background(), &demo.HiRequest{
		Name: "go",
		Msg:  "hi",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 0 {
		t.Fatalf("code: %v, msg: %s", res.Code, res.Msg)
	}
	t.Log(res.Data)
}

func handle(h http.Handler) {
	server := h.(*grpc.Server)
	demo.RegisterDemoServiceServer(server, &demo.DemoHandler{})
}
