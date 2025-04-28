package core

type Options struct {
	// 大小配置
	Size Size `json:"size,omitempty"`
	// 长连接
	Keepalive Keepalive `json:"keepalive,omitempty"`
}
