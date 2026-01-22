package config

import (
	"github.com/harluo/config"
	"github.com/harluo/grpc/internal/internal/core"
)

type Grpc struct {
	// 服务器端配置
	Server *core.Server `json:"server,omitempty"`
	// 客户端配置
	Clients map[string]core.Client `json:"clients,omitempty"`
	// gRPC配置
	Options core.Options `json:"options,omitempty"`
}

func newConfig(getter config.Getter) (grpc *Grpc, err error) {
	grpc = new(Grpc)
	err = getter.Get(&struct {
		Grpc *Grpc `json:"grpc,omitempty" validate:"required"`
	}{
		Grpc: grpc,
	})

	return
}
