package checker

import (
	"github.com/harluo/grpc/internal/test"
)

type Gateway interface {
	Handlers() []test.Handler
}
