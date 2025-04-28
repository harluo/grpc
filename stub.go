package grpc

import (
	"github.com/harluo/grpc/internal/kernel"
	"github.com/harluo/grpc/internal/stub"
	"google.golang.org/grpc"
)

type Stub = kernel.Stub

func NewStub[T any](fun func(grpc.ServiceRegistrar, T), server T) kernel.Stub {
	return stub.NewDefault(fun, server)
}
