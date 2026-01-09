package core

import (
	"fmt"
)

type Client struct {
	// 名称列表
	Names []string `json:"names,omitempty"`
	// 地址
	Host string `json:"host,omitempty" validate:"required,ip|hostname"`
	// 地址
	Hostname string `json:"hostname,omitempty" validate:"required,ip|hostname"`
	// 端口
	Port uint32 `default:"90" json:"port,omitempty" validate:"max=65535"`
}

func (c *Client) Addr() (addr string) {
	if c.Host != "" {
		addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	} else if c.Hostname != "" {
		addr = fmt.Sprintf("%s:%d", c.Hostname, c.Port)
	}

	return
}
