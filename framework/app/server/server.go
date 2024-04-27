package server

import "context"

const (
	KindGRPC = "grpc"
	KindHTTP = "http"
)

// Server is transport server.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
