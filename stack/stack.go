package stack

type Stack[T any] struct {
	head *node[T]
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (s *Stack[T]) Push(val T) {
	newNode := &node[T]{val: val, next: s.head}
	s.head = newNode
}

func (s *Stack[T]) Pop() (val T) {
	if s.Empty() {
		return val
	}

	head := s.head
	s.head = head.next
	head.next = nil
	return head.val
}

func (s *Stack[T]) Peek() (val T) {
	if s.Empty() {
		return val
	}
	return s.head.val
}

func (s *Stack[T]) Empty() bool {
	return s.head == nil
}
