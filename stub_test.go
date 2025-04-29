package grpc_test

import (
	"testing"

	"github.com/harluo/grpc"
	"github.com/harluo/grpc/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestNewStub(t *testing.T) {
	stub := grpc.NewStub(test.StubInt, 5)
	assert.NotNil(t, stub)
}
