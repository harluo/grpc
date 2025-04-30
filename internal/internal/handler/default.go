package handler

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/harluo/grpc/internal/internal/constant"
	"github.com/harluo/grpc/internal/internal/kernel"
	"github.com/harluo/grpc/internal/internal/param"
	"google.golang.org/grpc"
)

type Default struct {
	params *param.Handler
}

func NewDefault(params *param.Handler) *Default {
	return &Default{
		params: params,
	}
}

func (d *Default) Handle(
	ctx context.Context,
	gateway *runtime.ServeMux, conn *grpc.ClientConn,
	mux *http.ServeMux,
) (err error) {
	switch d.params.Type {
	case constant.HandlerTypeGrpc:
		err = d.params.Data.(kernel.Handler)(ctx, gateway, conn)
	case constant.HandlerTypeHttp:
		mux.Handle(d.params.Pattern, d.params.Data.(http.Handler))
	default:
		mux.Handle(d.params.Pattern, d.params.Data.(http.Handler))
	}

	return
}
