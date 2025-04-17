package internal

import (
	"github.com/harluo/grpc/internal/config/internal"
)

type Config struct {
	Server  *internal.Server
	Gateway *internal.Gateway
}

func NewConfig(server *internal.Server, gateway *internal.Gateway) *Config {
	return &Config{
		Server:  server,
		Gateway: gateway,
	}
}
