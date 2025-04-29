package test

import (
	"google.golang.org/grpc"
)

func StubInt(_ grpc.ServiceRegistrar, _ int) {}
