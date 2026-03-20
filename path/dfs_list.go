package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func dfsList(graph graph.AdjacencyList, curr, target int, visited set.Set[int], path *[]int) bool {
	if visited.Contains(curr) {
		return false
	}

	visited.Add(curr)

	*path = append(*path, curr)

	if curr == target {
		return true
	}

	for _, edge := range graph[curr] {
		if dfsList(graph, edge.To, target, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DepthFirstSearchList(graph graph.AdjacencyList, source, target int) []int {
	path := make([]int, 0, len(graph))
	dfsList(graph, source, target, set.New[int](), &path)
	return path
}
