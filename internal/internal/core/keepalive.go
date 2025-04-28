package core

import (
	"time"
)

type Keepalive struct {
	// 保持时长
	Time time.Duration `default:"10s" json:"time,omitempty"`
	// 超时
	Timeout time.Duration `default:"3s" json:"timeout,omitempty"`
	// 空闲时长
	Idle time.Duration `default:"3s" json:"idle,omitempty"`
	// 策略
	Policy KeepalivePolicy `json:"policy,omitempty"`
}
