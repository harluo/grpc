package kernel

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Unescape struct {
	// 模式
	Mode runtime.UnescapingMode `json:"mode,omitempty" validate:"max=3"`
}
