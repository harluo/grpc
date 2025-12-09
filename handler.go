package grpc

import (
	"github.com/harluo/grpc/internal/handler"
	"github.com/harluo/grpc/internal/kernel"
	"google.golang.org/grpc"
)

type Handler = kernel.Handler

func NewHandler[T any](fun func(grpc.ServiceRegistrar, T), server T) kernel.Handler {
	return handler.NewDefault(fun, server)
}
