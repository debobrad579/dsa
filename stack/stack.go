package stack

type Stack[T any] interface {
	Push(T)
	Pop() T
	Peek() T
	Empty() bool
}

type stack[T any] struct {
	head *node[T]
}

func New[T any]() Stack[T] {
	return &stack[T]{}
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (s *stack[T]) Push(val T) {
	newNode := &node[T]{val: val, next: s.head}
	s.head = newNode
}

func (s *stack[T]) Pop() (val T) {
	if s.Empty() {
		panic("stack is empty")
	}

	head := s.head
	s.head = head.next
	head.next = nil
	return head.val
}

func (s *stack[T]) Peek() (val T) {
	if s.Empty() {
		panic("stack is empty")
	}

	return s.head.val
}

func (s *stack[T]) Empty() bool {
	return s.head == nil
}
