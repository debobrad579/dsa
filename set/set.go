package set

type set[T comparable] map[T]struct{}

func New[T comparable](items ...T) set[T] {
	s := make(set[T])

	for _, item := range items {
		s.Add(item)
	}

	return s
}

func (s set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s set[T]) Remove(item T) {
	delete(s, item)
}

func (s set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s set[T]) Size() int {
	return len(s)
}
