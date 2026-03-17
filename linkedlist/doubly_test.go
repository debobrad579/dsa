package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/linkedlist"
)

func TestDoublyLength(t *testing.T) {
	var l linkedlist.Doubly[int]
	assert.Equal(t, 0, l.Length())
	l.Append(1)
	assert.Equal(t, 1, l.Length())
	l.Append(2)
	l.Append(3)
	assert.Equal(t, 3, l.Length())
}

func TestDoublyGet(t *testing.T) {
	var l linkedlist.Doubly[int]
	_, err := l.GetFromEnd(0)
	assert.Error(t, err)
	l.Append(10)
	l.Append(20)
	l.Append(30)
	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	val, err = l.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	_, err = l.Get(3)
	assert.Error(t, err)
	_, err = l.Get(-1)
	assert.Error(t, err)
}

func TestDoublyGetFromEnd(t *testing.T) {
	var l linkedlist.Doubly[int]
	_, err := l.GetFromEnd(0)
	assert.Error(t, err)
	l.Append(10)
	l.Append(20)
	l.Append(30)
	val, err := l.GetFromEnd(0)
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	val, err = l.GetFromEnd(2)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	_, err = l.GetFromEnd(3)
	assert.Error(t, err)
	_, err = l.GetFromEnd(-1)
	assert.Error(t, err)
}

func TestDoublyAppend(t *testing.T) {
	var l linkedlist.Doubly[int]
	l.Append(1)
	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	l.Append(2)
	l.Append(3)
	val, err = l.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
}

func TestDoublyPrepend(t *testing.T) {
	var l linkedlist.Doubly[int]
	l.Prepend(1)
	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	l.Prepend(2)
	val, err = l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
}

func TestDoublyInsert(t *testing.T) {
	var l linkedlist.Doubly[int]
	err := l.Insert(1, -1)
	assert.Error(t, err)
	err = l.Insert(1, 1)
	assert.Error(t, err)
	l.Append(10)
	l.Append(30)
	err = l.Insert(20, 1)
	assert.NoError(t, err)
	val, err := l.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	err = l.Insert(5, 0)
	assert.NoError(t, err)
	val, err = l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 5, val)
	err = l.Insert(99, 10)
	assert.Error(t, err)
}

func TestDoublyInsertFromEnd(t *testing.T) {
	var l linkedlist.Doubly[int]
	err := l.InsertFromEnd(1, -1)
	assert.Error(t, err)
	err = l.InsertFromEnd(1, 1)
	assert.Error(t, err)
	l.Append(10)
	l.Append(30)
	err = l.InsertFromEnd(20, 1)
	assert.NoError(t, err)
	val, err := l.GetFromEnd(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	err = l.InsertFromEnd(99, 10)
	assert.Error(t, err)
}

func TestDoublyDelete(t *testing.T) {
	var l linkedlist.Doubly[int]
	err := l.Delete(-1)
	assert.Error(t, err)
	err = l.Delete(0)
	assert.Error(t, err)
	l.Append(10)
	l.Append(20)
	l.Append(30)
	err = l.Delete(1)
	assert.NoError(t, err)
	val, err := l.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	err = l.Delete(0)
	assert.NoError(t, err)
	val, err = l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	err = l.Delete(5)
	assert.Error(t, err)
}

func TestDoublyDeleteFromEnd(t *testing.T) {
	var l linkedlist.Doubly[int]
	err := l.DeleteFromEnd(-1)
	assert.Error(t, err)
	err = l.DeleteFromEnd(0)
	assert.Error(t, err)
	l.Append(10)
	l.Append(20)
	l.Append(30)
	err = l.DeleteFromEnd(1)
	assert.NoError(t, err)
	val, err := l.GetFromEnd(1)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	err = l.DeleteFromEnd(0)
	assert.NoError(t, err)
	val, err = l.GetFromEnd(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	err = l.DeleteFromEnd(5)
	assert.Error(t, err)
}

func TestDoublyPop(t *testing.T) {
	var l linkedlist.Doubly[int]
	_, err := l.Pop()
	assert.Error(t, err)
	l.Append(10)
	l.Append(20)
	l.Append(30)
	val, err := l.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	assert.Equal(t, 2, l.Length())
	val, err = l.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	val, err = l.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	_, err = l.Pop()
	assert.Error(t, err)
}
