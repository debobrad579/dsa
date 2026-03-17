package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/stack"
)

func TestQueue(t *testing.T) {
	var s stack.Stack[int]
	assert.True(t, s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Peek())
}
