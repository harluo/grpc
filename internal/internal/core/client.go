package core

import (
	"fmt"
)

type Client struct {
	// 名称
	Name string `json:"name,omitempty" validate:"required_without=Names"`
	// 名称列表
	Names []string `json:"names,omitempty" validate:"required_without=Name"`
	// 地址
	Hostname string `json:"hostname,omitempty" validate:"required,ip|hostname"`
	// 端口
	Port uint32 `default:"90" json:"port,omitempty" validate:"max=65535"`
}

func (c *Client) Addr() string {
	return fmt.Sprintf("%s:%d", c.Hostname, c.Port)
}
