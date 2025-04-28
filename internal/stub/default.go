package stub

import (
	"google.golang.org/grpc"
)

type Default[T any] struct {
	fun    func(grpc.ServiceRegistrar, T)
	server T
}

func NewDefault[T any](fun func(grpc.ServiceRegistrar, T), server T) *Default[T] {
	return &Default[T]{
		fun:    fun,
		server: server,
	}
}

func (d *Default[T]) Register(server *grpc.Server) {
	d.fun(server, d.server)
}
