package core

import (
	"github.com/goexl/gox"
)

type Window struct {
	// 初始
	Initial gox.Bytes `default:"1GB" json:"initial,omitempty"`
	// 连接
	Connection gox.Bytes `default:"1GB" json:"connection,omitempty"`
}
