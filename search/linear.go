package search

func LinearSearch[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}
