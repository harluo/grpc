package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newServer,
		newClient,
		newGateway,
	).Build().Apply()
}
