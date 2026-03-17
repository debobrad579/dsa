package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/sort"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 5, 3, 7, 4, 0, 6, 8, 1, 2}
	expectedArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.QuickSort(arr)
	assert.Equal(t, expectedArr, arr)

	arr = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	expectedArr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	sort.QuickSort(arr)
	assert.Equal(t, expectedArr, arr)
}
