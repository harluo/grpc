package core

import (
	"time"
)

type Keepalive struct {
	Time    time.Duration   `default:"5m" json:"time,omitempty"`     // 保持时长
	Timeout time.Duration   `default:"20s" json:"timeout,omitempty"` // 超时
	Idle    time.Duration   `default:"3s" json:"idle,omitempty"`     // 空闲时长
	Policy  KeepalivePolicy `json:"policy,omitempty"`                // 策略
}
