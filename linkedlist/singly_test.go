package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/linkedlist"
)

func TestSinglyLength(t *testing.T) {
	var l linkedlist.Singly[int]
	assert.Equal(t, 0, l.Length())
	l.Append(1)
	assert.Equal(t, 1, l.Length())
	l.Append(2)
	l.Append(3)
	assert.Equal(t, 3, l.Length())
}

func TestSinglyGet(t *testing.T) {
	var l linkedlist.Singly[int]
	_, err := l.Get(0)
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

func TestSinglyAppend(t *testing.T) {
	var l linkedlist.Singly[int]
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

func TestSinglyPrepend(t *testing.T) {
	var l linkedlist.Singly[int]
	l.Prepend(1)
	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	l.Prepend(2)
	val, err = l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
}

func TestSinglyInsert(t *testing.T) {
	var l linkedlist.Singly[int]
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
}

func TestSinglyDelete(t *testing.T) {
	var l linkedlist.Singly[int]
	err := l.Delete(-1)
	assert.Error(t, err)
	err = l.Delete(0)
	assert.Error(t, err)
	l.Append(10)
	l.Append(30)
	err = l.Delete(1)
	assert.NoError(t, err)
	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
}
