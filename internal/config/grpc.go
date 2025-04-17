package config

import (
	"github.com/harluo/boot"
	"github.com/harluo/grpc/internal/config/internal"
)

type Grpc struct {
	// 服务器端配置
	Server *internal.Server `json:"server,omitempty"`
	// 网关配置
	Gateway *internal.Gateway `json:"gateway,omitempty"`
	// 客户端配置
	Clients []internal.Client `json:"clients,omitempty"`
	// gRPC配置
	Options internal.Options `json:"options,omitempty"`
}

func newConfig(config *boot.Config) (grpc *Grpc, err error) {
	grpc = new(Grpc)
	err = config.Build().Get(&struct {
		Grpc *Grpc `json:"grpc,omitempty" validate:"required"`
	}{
		Grpc: grpc,
	})

	return
}
