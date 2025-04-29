package checker

import (
	"github.com/harluo/grpc/internal/kernel"
)

type Gateway interface {
	Handlers() []kernel.Handler
}
