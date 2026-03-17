package sort

func SelectionSort(arr []int) {
	for i := range len(arr) {
		m := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		arr[i], arr[m] = arr[m], arr[i]
	}
}
