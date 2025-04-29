package param

import (
	"github.com/harluo/grpc/internal/internal/constant"
)

type Handler struct {
	Data    any
	Type    constant.HandlerType
	Pattern string
}

func NewHandler() *Handler {
	return new(Handler)
}
