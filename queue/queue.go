package queue

type Queue[T any] interface {
	Enqueue(T)
	Deque() T
	Peek() T
	Empty() bool
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() Queue[T] {
	return &queue[T]{}
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (q *queue[T]) Enqueue(val T) {
	newNode := &node[T]{val: val}

	if q.Empty() {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}

	q.tail = newNode
}

func (q *queue[T]) Deque() (val T) {
	if q.Empty() {
		panic("queue is empty")
	}

	head := q.head
	q.head = head.next

	if q.Empty() {
		q.tail = nil
	}

	head.next = nil
	return head.val
}

func (q *queue[T]) Peek() (val T) {
	if q.Empty() {
		panic("queue is empty")
	}

	return q.head.val
}

func (q *queue[T]) Empty() bool {
	return q.head == nil
}
