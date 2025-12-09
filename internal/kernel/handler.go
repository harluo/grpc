package kernel

import (
	"google.golang.org/grpc"
)

type Handler interface {
	Handle(*grpc.Server)
}
