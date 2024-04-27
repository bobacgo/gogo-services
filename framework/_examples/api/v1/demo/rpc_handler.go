package demo

import "context"

var _ DemoServiceServer = (*DemoHandler)(nil)

type DemoHandler struct{}

func (h DemoHandler) Ping(ctx context.Context, in *PingRequest) (*DemoResponse, error) {
	addr := in.GetAddr()
	return &DemoResponse{
		Code: 0,
		Msg:  "ok",
		Data: addr,
	}, nil
}

func (h DemoHandler) Hi(ctx context.Context, in *HiRequest) (*DemoResponse, error) {
	return &DemoResponse{
		Code: 0,
		Msg:  "ok",
		Data: in.GetName() + in.GetMsg(),
	}, nil
}
