package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func dfsMatrix(g graph.AdjacencyMatrix, curr, dest int, visited set.Set[int], path *[]int) bool {
	if visited.Contains(curr) {
		return false
	}

	visited.Add(curr)

	*path = append(*path, curr)

	if curr == dest {
		return true
	}

	for i := range len(g[curr]) {
		if g[curr][i] != 0 && dfsMatrix(g, i, dest, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DepthFirstSearchMatrix(g graph.AdjacencyMatrix, source, dest int) []int {
	path := make([]int, 0, len(g))
	dfsMatrix(g, source, dest, set.New[int](), &path)
	return path
}
