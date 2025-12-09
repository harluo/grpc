package grpc_test

import (
	"testing"

	"github.com/harluo/grpc"
	"github.com/harluo/grpc/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	handler := grpc.NewHandler(test.HandlerInt, 5)
	assert.NotNil(t, handler)
}
