package kernel

import (
	"time"
)

type Timeout struct {
	// 读
	Read time.Duration `default:"15s" json:"read,omitempty"`
	// 头
	Header time.Duration `default:"15s" json:"header,omitempty"`
}
