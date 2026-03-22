package path

import (
	"errors"
	"math"
	"slices"

	"github.com/debobrad579/dsa/graph"
)

func BellmanFord(g graph.AdjacencyList, source, dest int) ([]int, error) {
	prev := make([]int, len(g))
	distances := make([]float64, len(g))
	for i := range len(g) {
		prev[i] = -1
		distances[i] = math.Inf(1)
	}
	distances[source] = 0

	for range len(g) - 1 {
		for i := range g {
			for _, v := range g[i] {
				distance := distances[i] + float64(v.Weight)

				if distance < distances[v.To] {
					distances[v.To] = distance
					prev[v.To] = i
				}
			}
		}
	}

	for i := range g {
		for _, v := range g[i] {
			if distances[i]+float64(v.Weight) < distances[v.To] {
				return nil, errors.New("negative cycle found")
			}
		}
	}

	if math.IsInf(distances[dest], 1) {
		return nil, errors.New("destination unreachable")
	}

	path := make([]int, 0)
	for curr := dest; curr != -1; curr = prev[curr] {
		path = append(path, curr)
	}

	slices.Reverse(path)

	return path, nil
}
