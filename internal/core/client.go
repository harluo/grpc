package core

import (
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/harluo/grpc/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// Client gRPC客户端封装
type Client struct {
	connections map[string]*grpc.ClientConn

	_ gox.Pointerized
}

func newClient(gc *config.Grpc) (client *Client, err error) {
	client = new(Client) // 避免空指针错误
	if len(gc.Clients) == 0 {
		err = exception.New().Message("缺乏客户端配置").Build()
	}
	if nil != err {
		return
	}

	opts := make([]grpc.DialOption, 0, 8)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithInitialWindowSize(int32(gc.Options.Size.Window.Initial)))
	opts = append(opts, grpc.WithInitialConnWindowSize(int32(gc.Options.Size.Window.Connection)))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(int(gc.Options.Size.Msg.Send))))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(gc.Options.Size.Msg.Receive))))
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                gc.Options.Keepalive.Time,
		Timeout:             gc.Options.Keepalive.Timeout,
		PermitWithoutStream: gc.Options.Keepalive.Policy.Permit,
	}))

	connections := make(map[string]*grpc.ClientConn)
	for name, conf := range gc.Clients {
		var connection *grpc.ClientConn
		if connection, err = grpc.Dial(conf.Addr(), opts...); nil != err {
			return
		}

		if "" != name {
			connections[name] = connection
		}
		for _, _name := range conf.Names {
			connections[_name] = connection
		}
	}
	client.connections = connections

	return
}

func (c *Client) Connection(name string) *grpc.ClientConn {
	return c.connections[name]
}
