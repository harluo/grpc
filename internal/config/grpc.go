package config

import (
	"github.com/harluo/config"
	"github.com/harluo/grpc/internal/internal/kernel"
)

type Grpc struct {
	// 服务器端配置
	Server *kernel.Server `json:"server,omitempty"`
	// 网关配置
	Gateway *kernel.Gateway `json:"gateway,omitempty"`
	// 客户端配置
	Clients []kernel.Client `json:"clients,omitempty"`
	// gRPC配置
	Options kernel.Options `json:"options,omitempty"`
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
