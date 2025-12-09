package test

import (
	"google.golang.org/grpc"
)

func HandlerInt(_ grpc.ServiceRegistrar, _ int) {}
