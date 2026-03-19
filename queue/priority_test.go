package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/queue"
)

func setupPriorityQueue() queue.PriorityQueue[int] {
	return queue.NewPriorityQueue[int](func(a, b int) bool { return a < b })
}

func TestPriorityQueueInsertSingle(t *testing.T) {
	pq := setupPriorityQueue()

	pq.Push(10)

	assert.Equal(t, 1, pq.Length())
	assert.Equal(t, 10, pq.Pop())
}

func TestPriorityQueueMinHeapProperty(t *testing.T) {
	pq := setupPriorityQueue()

	pq.Push(5)
	pq.Push(3)
	pq.Push(8)
	pq.Push(1)

	assert.Equal(t, 1, pq.Peek())
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 3, pq.Peek())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 5, pq.Peek())
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 8, pq.Peek())
	assert.Equal(t, 8, pq.Pop())
}

func TestPriorityQueueDelete(t *testing.T) {
	pq := setupPriorityQueue()

	pq.Push(10)
	pq.Push(4)
	pq.Push(7)

	assert.Equal(t, 4, pq.Pop())

	pq.Push(2)
	pq.Push(6)

	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 6, pq.Pop())
	assert.Equal(t, 7, pq.Pop())
	assert.Equal(t, 10, pq.Pop())
}

func TestPriorityQueueDuplicates(t *testing.T) {
	pq := setupPriorityQueue()

	pq.Push(5)
	pq.Push(1)
	pq.Push(1)
	pq.Push(3)

	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 5, pq.Pop())
}
