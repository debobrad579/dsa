package sort

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func merge(a, b []int) []int {
	final := make([]int, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			final[i+j] = a[i]
			i++
		} else {
			final[i+j] = b[j]
			j++
		}
	}

	for ; i < len(a); i++ {
		final[i+j] = a[i]
	}

	for ; j < len(b); j++ {
		final[i+j] = b[j]
	}

	return final
}
