package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/search"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{3, 2, 5, 1, 4}
	assert.Equal(t, 0, search.LinearSearch(arr, 3))
	assert.Equal(t, 2, search.LinearSearch(arr, 5))
	assert.Equal(t, -1, search.LinearSearch(arr, 6))
}
