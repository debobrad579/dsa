package sort

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) {
	for i := range len(arr) {
		for j := range len(arr) - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
