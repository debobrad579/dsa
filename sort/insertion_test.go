package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/sort"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{9, 5, 3, 7, 4, 0, 6, 8, 1, 2}
	expectedArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.InsertionSort(arr)
	assert.Equal(t, expectedArr, arr)

	arr = []int{0, 1, 2, 3, 5, 4, 6, 8, 9, 7, 15, 10, 11, 13, 12, 14}
	expectedArr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	sort.InsertionSort(arr)
	assert.Equal(t, expectedArr, arr)
}
