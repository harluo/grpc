package core

import (
	"net"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/harluo/grpc/internal/kernel"
)

func (s *Server) setupGrpc(register kernel.Register, listener net.Listener) (err error) {
	for _, stub := range register.Stubs() {
		stub.Register(s.rpc)
	}
	fields := gox.Fields[any]{
		field.New("name", s.server.Name),
		field.New("addr", s.server.Addr()),
	}
	s.logger.Info("启动服务成功", fields[0], fields[1:]...)
	if !s.gatewayEnabled() || (s.gatewayEnabled() && s.diffPort()) {
		s.started = true
		s.wait.Add(1)
		go s.serveRpc(listener, &fields)
	}

	return
}

func (s *Server) serveRpc(listener net.Listener, fields *gox.Fields[any]) {
	defer s.wait.Done()

	if err := s.rpc.Serve(listener); nil != err {
		s.logger.Error("启动服务出错", field.Error(err), *fields...)
	}
}
