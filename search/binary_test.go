package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/search"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, 0, search.BinarySearch(arr, 1))
	assert.Equal(t, 3, search.BinarySearch(arr, 4))
	assert.Equal(t, 7, search.BinarySearch(arr, 8))
	assert.Equal(t, -1, search.BinarySearch(arr, 9))

	arr = []int{1, 3, 10, 14, 21, 22, 25, 29}
	assert.Equal(t, 0, search.BinarySearch(arr, 1))
	assert.Equal(t, 2, search.BinarySearch(arr, 10))
	assert.Equal(t, 5, search.BinarySearch(arr, 22))
	assert.Equal(t, -1, search.BinarySearch(arr, 9))
}
