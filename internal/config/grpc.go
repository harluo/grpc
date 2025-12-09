package config

import (
	"github.com/harluo/config"
	"github.com/harluo/grpc/internal/internal/core"
)

type Grpc struct {
	// 服务器端配置
	Server *core.Server `json:"server,omitempty"`
	// 客户端配置
	Clients []core.Client `json:"clients,omitempty"`
	// gRPC配置
	Options core.Options `json:"options,omitempty"`
}

func newConfig(config config.Getter) (grpc *Grpc, err error) {
	grpc = new(Grpc)
	err = config.Get(&struct {
		Grpc *Grpc `json:"grpc,omitempty" validate:"required"`
	}{
		Grpc: grpc,
	})

	return
}
