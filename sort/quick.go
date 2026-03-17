package sort

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, lo, hi int) {
	if lo < hi {
		mid := partition(arr, lo, hi)
		qs(arr, lo, mid-1)
		qs(arr, mid+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	mid := lo + (hi-lo)/2
	pivot := arr[mid]
	arr[mid], arr[hi] = arr[hi], arr[mid]

	i := lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	return i + 1
}
