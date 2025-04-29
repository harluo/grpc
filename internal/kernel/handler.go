package kernel

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Handler interface {
	Handle(context.Context, *runtime.ServeMux, *grpc.ClientConn, *http.ServeMux) error
}
