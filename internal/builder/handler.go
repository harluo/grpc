package builder

import (
	"net/http"

	"github.com/harluo/grpc/internal/internal/constant"
	"github.com/harluo/grpc/internal/internal/handler"
	"github.com/harluo/grpc/internal/internal/kernel"
	"github.com/harluo/grpc/internal/internal/param"
)

type Handler struct {
	params *param.Handler
}

func NewHandler() *Handler {
	return &Handler{
		params: param.NewHandler(),
	}
}

func (h *Handler) Http(pattern string, handler *http.Handler) (self *Handler) {
	h.params.Type = constant.HandlerTypeHttp
	h.params.Pattern = pattern
	h.params.Data = handler

	return
}

func (h *Handler) Grpc(handler kernel.Handler) (self *Handler) {
	h.params.Type = constant.HandlerTypeGrpc
	h.params.Data = handler

	return
}

func (h *Handler) Build() *handler.Default {
	return handler.NewDefault(h.params)
}
