package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func dfsList(g graph.AdjacencyList, curr, dest int, visited set.Set[int], path *[]int) bool {
	if visited.Contains(curr) {
		return false
	}

	visited.Add(curr)

	*path = append(*path, curr)

	if curr == dest {
		return true
	}

	for _, edge := range g[curr] {
		if dfsList(g, edge.To, dest, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DepthFirstSearchList(g graph.AdjacencyList, source, dest int) []int {
	path := make([]int, 0, len(g))
	dfsList(g, source, dest, set.New[int](), &path)
	return path
}
