package lru_test

import (
	"testing"

	"github.com/debobrad579/dsa/lru"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	c := lru.New[string, int](3)
	c.Update("a", 1)
	c.Update("b", 2)

	val, ok := c.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	_, ok = c.Get("z")
	assert.False(t, ok, "missing key should return false")
}

func TestUpdate(t *testing.T) {
	c := lru.New[string, int](3)
	c.Update("a", 1)
	c.Update("a", 99)

	val, ok := c.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 99, val, "update should overwrite existing value")
}

func TestEviction(t *testing.T) {
	c := lru.New[string, int](3)
	c.Update("a", 1)
	c.Update("b", 2)
	c.Update("c", 3)
	c.Update("d", 4)

	_, ok := c.Get("a")
	assert.False(t, ok, "least recently used key should be evicted")
	_, ok = c.Get("b")
	assert.True(t, ok)
	_, ok = c.Get("c")
	assert.True(t, ok)
	_, ok = c.Get("d")
	assert.True(t, ok)
}

func TestGetPromotesRecency(t *testing.T) {
	c := lru.New[string, int](3)
	c.Update("a", 1)
	c.Update("b", 2)
	c.Update("c", 3)
	c.Get("a")
	c.Update("d", 4)

	_, ok := c.Get("b")
	assert.False(t, ok, "least recently used key should be evicted after get promotes another")
	_, ok = c.Get("a")
	assert.True(t, ok)
}

func TestUpdatePromotesRecency(t *testing.T) {
	c := lru.New[string, int](3)
	c.Update("a", 1)
	c.Update("b", 2)
	c.Update("c", 3)
	c.Update("a", 10)
	c.Update("d", 4)

	_, ok := c.Get("b")
	assert.False(t, ok, "least recently used key should be evicted after update promotes another")
	_, ok = c.Get("a")
	assert.True(t, ok)
	val, _ := c.Get("a")
	assert.Equal(t, 10, val)
}

func TestCapacityOne(t *testing.T) {
	c := lru.New[string, int](1)
	c.Update("a", 1)
	c.Update("b", 2)

	_, ok := c.Get("a")
	assert.False(t, ok, "only one item should fit in cache")
	val, ok := c.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
}
