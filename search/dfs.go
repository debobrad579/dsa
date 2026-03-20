package search

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func dfs(graph graph.AdjacencyList, curr, target int, visited set.Set[int], path *[]int) bool {
	if visited.Contains(curr) {
		return false
	}

	visited.Add(curr)

	*path = append(*path, curr)

	if curr == target {
		return true
	}

	for _, edge := range graph[curr] {
		if dfs(graph, edge.To, target, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func DepthFirstSearch(graph graph.AdjacencyList, source, target int) []int {
	path := make([]int, 0, len(graph))
	dfs(graph, source, target, set.New[int](), &path)
	return path
}
