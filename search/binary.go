package search

func BinarySearch(arr []int, target int) int {
	return search(arr, 0, len(arr)-1, target)
}

func search(arr []int, lo, hi, target int) int {
	for {
		mid := lo + (hi-lo)/2
		v := arr[mid]

		if v == target {
			return mid
		}

		if lo >= hi {
			break
		}

		if v > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}
