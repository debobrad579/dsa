package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func dfsMatrix(graph graph.AdjacencyMatrix, curr, target int, visited set.Set[int], path *[]int) bool {
	if visited.Contains(curr) {
		return false
	}

	visited.Add(curr)

	*path = append(*path, curr)

	if curr == target {
		return true
	}

	for i := range len(graph[curr]) {
		if graph[curr][i] != 0 && dfsMatrix(graph, i, target, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DepthFirstSearchMatrix(graph graph.AdjacencyMatrix, source, target int) []int {
	path := make([]int, 0, len(graph))
	dfsMatrix(graph, source, target, set.New[int](), &path)
	return path
}
