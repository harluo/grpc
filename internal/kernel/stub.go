package kernel

import (
	"google.golang.org/grpc"
)

type Stub interface {
	Register(*grpc.Server)
}
