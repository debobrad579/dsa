package queue

type PriorityQueue[T any] interface {
	Push(T)
	Pop() T
	Peek() T
	Length() int
}

type priorityQueue[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) PriorityQueue[T] {
	return &priorityQueue[T]{less: less}
}

func (pq *priorityQueue[T]) parentIdx(i int) int {
	return (i - 1) / 2
}

func (pq *priorityQueue[T]) leftIdx(i int) int {
	return 2*i + 1
}

func (pq *priorityQueue[T]) rightIdx(i int) int {
	return 2*i + 2
}

func (pq *priorityQueue[T]) Length() int {
	return len(pq.data)
}

func (pq *priorityQueue[T]) Peek() T {
	if len(pq.data) == 0 {
		panic("priority queue is empty")
	}

	return pq.data[0]
}

func (pq *priorityQueue[T]) heapifyUp(i int) {
	for i > 0 {
		parentIdx := pq.parentIdx(i)

		if !pq.less(pq.data[i], pq.data[parentIdx]) {
			break
		}

		pq.data[i], pq.data[parentIdx] = pq.data[parentIdx], pq.data[i]
		i = parentIdx
	}
}

func (pq *priorityQueue[T]) heapifyDown(i int) {
	for {
		leftIdx := pq.leftIdx(i)
		rightIdx := pq.rightIdx(i)
		smallestIdx := i

		if leftIdx < len(pq.data) && pq.less(pq.data[leftIdx], pq.data[smallestIdx]) {
			smallestIdx = leftIdx
		}
		if rightIdx < len(pq.data) && pq.less(pq.data[rightIdx], pq.data[smallestIdx]) {
			smallestIdx = rightIdx
		}

		if smallestIdx == i {
			break
		}

		pq.data[i], pq.data[smallestIdx] = pq.data[smallestIdx], pq.data[i]
		i = smallestIdx
	}
}

func (pq *priorityQueue[T]) Push(val T) {
	pq.data = append(pq.data, val)
	pq.heapifyUp(len(pq.data) - 1)
}

func (pq *priorityQueue[T]) Pop() T {
	if len(pq.data) == 0 {
		panic("priority queue is empty")
	}

	out := pq.data[0]
	pq.data[0] = pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]
	pq.heapifyDown(0)
	return out
}
