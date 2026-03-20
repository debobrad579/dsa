package path

import (
	"math"
	"slices"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/set"
)

func getLowestUnvisited(visited set.Set[int], distances []float64) int {
	idx := -1
	lowestDistance := math.Inf(1)

	for i := range distances {
		if visited.Contains(i) {
			continue
		}

		if distances[i] < lowestDistance {
			idx = i
			lowestDistance = distances[i]
		}
	}

	return idx
}

func DijkstraShortestPath(graph graph.AdjacencyList, source, target int) []int {
	visited := set.New[int]()
	prev := make([]int, len(graph))
	distances := make([]float64, len(graph))

	for i := range len(graph) {
		prev[i] = -1
		distances[i] = math.Inf(1)
	}

	distances[source] = 0

	for curr := getLowestUnvisited(visited, distances); curr != -1; curr = getLowestUnvisited(visited, distances) {
		visited.Add(curr)
		if curr == target {
			break
		}

		for _, edge := range graph[curr] {
			if visited.Contains(edge.To) {
				continue
			}

			if dist := distances[curr] + float64(edge.Weight); dist < distances[edge.To] {
				prev[edge.To] = curr
				distances[edge.To] = dist
			}
		}
	}

	if distances[target] == math.Inf(1) {
		return []int{}
	}

	path := make([]int, 0)
	for curr := target; curr != -1; curr = prev[curr] {
		path = append(path, curr)
	}

	slices.Reverse(path)

	return path
}
