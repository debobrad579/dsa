package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()
	assert.True(t, q.Empty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 1, q.Deque())
	assert.Equal(t, 2, q.Peek())
}
