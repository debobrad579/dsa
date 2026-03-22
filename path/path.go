package path

import "slices"

func reconstructPath(prev []int, source, dest int) []int {
	path := make([]int, 0)

	for curr := dest; curr != -1; curr = prev[curr] {
		path = append(path, curr)
	}

	slices.Reverse(path)

	if len(path) == 0 || path[0] != source {
		return []int{}
	}

	return path
}
