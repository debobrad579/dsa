package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/sort"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 5, 3, 7, 4, 1, 6, 8}
	expectedArr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sort.BubbleSort(arr)
	assert.Equal(t, expectedArr, arr)

	arr = []int{5, 4, 3, 2, 1}
	expectedArr = []int{1, 2, 3, 4, 5}
	sort.BubbleSort(arr)
	assert.Equal(t, expectedArr, arr)
}
