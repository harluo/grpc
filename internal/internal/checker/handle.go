package checker

import (
	"net/http"
)

type Handle interface {
	Handles() []http.Handler
}
