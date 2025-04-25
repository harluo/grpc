package kernel

import (
	"github.com/goexl/gox"
)

type Msg struct {
	// 发送大小
	// 4GB
	Send gox.Bytes `default:"4GB" json:"send,omitempty"`
	// 接收大小
	// 4GB
	Receive gox.Bytes `default:"4GB" json:"receive,omitempty"`
}
