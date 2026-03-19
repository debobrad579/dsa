package set_test

import (
	"testing"

	"github.com/debobrad579/dsa/set"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := set.New[int](1, 2, 3, 3)

	assert.Equal(t, 3, s.Size(), "should not include duplicates")
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestAdd(t *testing.T) {
	s := set.New[int]()

	s.Add(1)
	s.Add(2)
	s.Add(2)

	assert.Equal(t, 2, s.Size(), "adding duplicate should not increase size")
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
}

func TestRemove(t *testing.T) {
	s := set.New[int](1, 2, 3)

	s.Remove(2)

	assert.Equal(t, 2, s.Size())
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(3))
}

func TestSize(t *testing.T) {
	s := set.New[int]()

	assert.Equal(t, 0, s.Size())

	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	s.Remove(1)

	assert.Equal(t, 1, s.Size())
}
