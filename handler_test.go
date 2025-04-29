package grpc_test

import (
	"testing"

	"github.com/harluo/grpc"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	handler := grpc.NewHandler().Build()
	assert.NotNil(t, handler)
}
