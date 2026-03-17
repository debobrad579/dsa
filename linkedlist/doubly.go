package linkedlist

import "errors"

type Doubly[T any] struct {
	head *nodeDoubly[T]
	tail *nodeDoubly[T]
}

type nodeDoubly[T any] struct {
	val  T
	next *nodeDoubly[T]
	prev *nodeDoubly[T]
}

func (l *Doubly[T]) Length() (i int) {
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

func (l *Doubly[T]) Insert(val T, i int) (err error) {
	if i == 0 {
		l.Prepend(val)
		return
	}

	if i < 0 || l.head == nil {
		return errors.New("index out of range")
	}

	currentNode := l.head
	for range i - 1 {
		if currentNode.next == nil {
			return errors.New("index out of range")
		}
		currentNode = currentNode.next
	}

	newNode := &nodeDoubly[T]{val: val, next: currentNode.next, prev: currentNode}
	if currentNode.next != nil {
		currentNode.next.prev = newNode
	} else {
		l.tail = newNode
	}

	currentNode.next = newNode
	return nil
}

func (l *Doubly[T]) InsertFromEnd(val T, i int) (err error) {
	if i == 0 {
		l.Append(val)
		return
	}

	if i < 0 || l.tail == nil {
		return errors.New("index out of range")
	}

	currentNode := l.tail
	for range i - 1 {
		if currentNode.prev == nil {
			return errors.New("index out of range")
		}
		currentNode = currentNode.prev
	}

	newNode := &nodeDoubly[T]{val: val, next: currentNode, prev: currentNode.prev}
	if currentNode.prev != nil {
		currentNode.prev.next = newNode
	} else {
		l.head = newNode
	}

	currentNode.prev = newNode
	return nil
}

func (l *Doubly[T]) Delete(i int) (err error) {
	if l.head == nil {
		return errors.New("index out of range")
	}

	if i == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		return
	}

	currentNode := l.head
	for range i - 1 {
		if currentNode.next == nil {
			return errors.New("index out of range")
		}
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		return errors.New("index out of range")
	}

	currentNode.next = currentNode.next.next
	if currentNode.next != nil {
		currentNode.next.prev = currentNode
	} else {
		l.tail = currentNode
	}

	return nil
}

func (l *Doubly[T]) DeleteFromEnd(i int) (err error) {
	if l.tail == nil {
		return errors.New("index out of range")
	}

	if i == 0 {
		l.tail = l.tail.prev
		if l.tail != nil {
			l.tail.next = nil
		} else {
			l.head = nil
		}
		return
	}

	currentNode := l.tail
	for range i - 1 {
		if currentNode.prev == nil {
			return errors.New("index out of range")
		}
		currentNode = currentNode.prev
	}

	if currentNode.prev == nil {
		return errors.New("index out of range")
	}

	currentNode.prev = currentNode.prev.prev
	if currentNode.prev != nil {
		currentNode.prev.next = currentNode
	} else {
		l.head = currentNode
	}

	return nil
}

func (l *Doubly[T]) Append(val T) {
	newNode := &nodeDoubly[T]{val: val, prev: l.tail}
	if l.tail != nil {
		l.tail.next = newNode
	}

	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
}

func (l *Doubly[T]) Prepend(val T) {
	newNode := &nodeDoubly[T]{val: val, next: l.head}
	if l.head != nil {
		l.head.prev = newNode
	}

	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
}

func (l *Doubly[T]) Get(i int) (val T, err error) {
	if i < 0 || l.head == nil {
		return val, errors.New("index out of range")
	}

	currentNode := l.head
	for range i {
		if currentNode.next == nil {
			return val, errors.New("index out of range")
		}

		currentNode = currentNode.next
	}

	return currentNode.val, nil
}

func (l *Doubly[T]) GetFromEnd(i int) (val T, err error) {
	if i < 0 || l.tail == nil {
		return val, errors.New("index out of range")
	}

	currentNode := l.tail
	for range i {
		if currentNode.prev == nil {
			return val, errors.New("index out of range")
		}

		currentNode = currentNode.prev
	}

	return currentNode.val, nil
}

func (l *Doubly[T]) Pop() (val T, err error) {
	if l.tail == nil {
		return val, errors.New("linked list is empty")
	}

	val = l.tail.val
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return val, nil
}
