package core

import (
	"context"
	"net"
	"net/http"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/grpc/internal/config"
	"github.com/harluo/grpc/internal/internal/constant"
	"github.com/harluo/grpc/internal/internal/core"
	"github.com/harluo/grpc/internal/kernel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// Server gRPC服务器封装
type Server struct {
	rpc    *grpc.Server
	config *core.Server
	logger log.Logger

	_ gox.Pointerized
}

func newServer(config *config.Grpc, logger log.Logger) (server *Server, mux *http.ServeMux, err error) {
	server = new(Server)
	if nil == config.Server {
		err = exception.New().Message("缺乏服务器配置").Build()
	}
	if nil != err {
		return
	}

	opts := make([]grpc.ServerOption, 0, 8)
	opts = append(opts, grpc.InitialWindowSize(int32(config.Options.Size.Window.Initial)))
	opts = append(opts, grpc.InitialConnWindowSize(int32(config.Options.Size.Window.Connection)))
	opts = append(opts, grpc.MaxSendMsgSize(int(config.Options.Size.Msg.Send)))
	opts = append(opts, grpc.MaxRecvMsgSize(int(config.Options.Size.Msg.Receive)))
	opts = append(opts, grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: config.Options.Keepalive.Policy.Permit,
	}))
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: config.Options.Keepalive.Idle,
		Time:              config.Options.Keepalive.Time,
		Timeout:           config.Options.Keepalive.Timeout,
	}))

	server.rpc = grpc.NewServer(opts...)
	server.config = config.Server
	server.logger = logger

	return
}

func (s *Server) Start(_ context.Context, router kernel.Router) (err error) {
	if *s.config.Reflection { // 反射，在gRPC接口调试时，可以反射出方法和参数
		reflection.Register(s.rpc)
	}
	for _, handler := range router.Handlers() {
		handler.Handle(s.rpc)
	}
	if rpc, le := net.Listen(constant.Tcp, s.config.Addr()); nil != le {
		err = le
		s.logger.Error("监控端口出错", field.Error(err), field.New("addr", s.config.Addr()))
	} else if se := s.rpc.Serve(rpc); nil != se {
		err = se
		s.logger.Error("启动服务出错", field.Error(err))
	}

	return
}

func (s *Server) Stop(_ context.Context) (err error) {
	s.rpc.GracefulStop()

	return
}
