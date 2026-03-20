package path

import (
	"math"
	"slices"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/queue"
)

func DijkstraShortestPath(g graph.AdjacencyList, source, target int) []int {
	prev := make([]int, len(g))
	distances := make([]float64, len(g))
	for i := range len(g) {
		prev[i] = -1
		distances[i] = math.Inf(1)
	}
	distances[source] = 0

	pq := queue.NewPriorityQueue(func(a, b graph.Edge) bool { return a.Weight < b.Weight })
	pq.Push(graph.Edge{To: source, Weight: 0})

	for pq.Length() != 0 {
		curr := pq.Pop()
		if curr.Weight != int(distances[curr.To]) {
			continue
		}

		if curr.To == target {
			break
		}

		for _, edge := range g[curr.To] {
			if dist := distances[curr.To] + float64(edge.Weight); dist < distances[edge.To] {
				prev[edge.To] = curr.To
				distances[edge.To] = dist
				pq.Push(graph.Edge{To: edge.To, Weight: int(dist)})
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
