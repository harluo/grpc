package core

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/harluo/grpc/internal/config"
	"github.com/harluo/grpc/internal/internal/constant"
	"github.com/harluo/grpc/internal/internal/kernel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// Server gRPC服务器封装
type Server struct {
	rpc  *grpc.Server
	http *http.Server
	mux  *http.ServeMux

	server  *kernel.Server
	gateway *kernel.Gateway

	wait    *sync.WaitGroup
	started bool
	logger  log.Logger

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

	mux = http.NewServeMux()
	server.rpc = grpc.NewServer(opts...)
	server.mux = mux
	server.logger = logger

	return
}

func (s *Server) Serve(register Register) (err error) {
	if *s.server.Reflection { // 反射，在gRPC接口调试时，可以反射出方法和参数
		reflection.Register(s.rpc)
	}

	if rpc, gateway, le := s.listeners(); nil != le {
		err = le
	} else if gre := s.setupGrpc(register, rpc); nil != gre {
		err = gre
	} else if gwe := s.setupGateway(register, gateway); nil != gwe {
		err = gwe
	}
	s.wait.Wait()

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	s.rpc.GracefulStop()
	if nil != s.http {
		err = s.http.Shutdown(ctx)
	}

	return
}

func (s *Server) diffPort() bool {
	return s.gateway.Port != s.server.Port
}

func (s *Server) listeners() (rpc net.Listener, gateway net.Listener, err error) {
	if listener, re := net.Listen(constant.Tcp, s.server.Addr()); nil != re { // gRPC端口必须监听
		err = re
	} else if s.gatewayEnabled() && s.diffPort() { // 如果网关开启且端口不一样
		rpc = listener
		gateway, err = net.Listen(constant.Tcp, s.gateway.Addr())
	} else { // 其它情况，监听端口都是一样的
		rpc = listener
		gateway = listener
	}

	return
}

func (s *Server) gatewayEnabled() bool {
	return nil != s.gateway && s.gateway.Enable()
}
