package queue

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (q *Queue[T]) Enqueue(val T) {
	newNode := &node[T]{val: val}

	if q.Empty() {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}

	q.tail = newNode
}

func (q *Queue[T]) Deque() (val T) {
	if q.Empty() {
		return val
	}

	head := q.head
	q.head = head.next

	if q.Empty() {
		q.tail = nil
	}

	head.next = nil
	return head.val
}

func (q *Queue[T]) Peek() (val T) {
	if q.Empty() {
		return val
	}
	return q.head.val
}

func (q *Queue[T]) Empty() bool {
	return q.head == nil
}
