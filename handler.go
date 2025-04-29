package grpc

import (
	"github.com/harluo/grpc/internal/builder"
	"github.com/harluo/grpc/internal/kernel"
)

// Handler 注册器
type Handler = kernel.Handler

func NewHandler() *builder.Handler {
	return builder.NewHandler()
}
