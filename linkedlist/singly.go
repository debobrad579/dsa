package linkedlist

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

type Singly[T any] struct {
	head *node[T]
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (l *Singly[T]) Length() (i int) {
	if l.head == nil {
		return 0
	}

	i++

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
		i++
	}

	return i
}

func (l *Singly[T]) Insert(val T, i int) (err error) {
	if i == 0 {
		l.Prepend(val)
		return
	}

	if i < 0 || l.head == nil {
		return ErrIndexOutOfRange
	}

	currentNode := l.head
	for range i - 1 {
		if currentNode.next == nil {
			return ErrIndexOutOfRange
		}

		currentNode = currentNode.next
	}

	currentNode.next = &node[T]{val: val, next: currentNode.next}
	return nil
}

func (l *Singly[T]) Delete(i int) (err error) {
	if l.head == nil {
		return ErrIndexOutOfRange
	}

	if i == 0 {
		l.head = l.head.next
		return
	}

	currentNode := l.head
	for range i - 1 {
		if currentNode.next == nil {
			return ErrIndexOutOfRange
		}

		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		return ErrIndexOutOfRange
	}

	currentNode.next = currentNode.next.next
	return nil
}

func (l *Singly[T]) Append(val T) {
	if l.head == nil {
		l.head = &node[T]{val: val}
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &node[T]{val: val}
}

func (l *Singly[T]) Prepend(val T) {
	l.head = &node[T]{val: val, next: l.head}
}

func (l *Singly[T]) Get(i int) (val T, err error) {
	if i < 0 || l.head == nil {
		return val, ErrIndexOutOfRange
	}

	currentNode := l.head
	for range i {
		if currentNode.next == nil {
			return val, ErrIndexOutOfRange
		}

		currentNode = currentNode.next
	}

	return currentNode.val, nil
}
