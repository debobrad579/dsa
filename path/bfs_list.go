package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/queue"
	"github.com/debobrad579/dsa/set"
)

func BreadthFirstSearchList(g graph.AdjacencyList, source, dest int) []int {
	q := queue.New[int]()
	q.Enqueue(source)
	visited := set.New(source)
	prev := make([]int, len(g))
	for i := range len(g) {
		prev[i] = -1
	}

	for !q.Empty() {
		curr := q.Deque()
		if curr == dest {
			break
		}

		for _, edge := range g[curr] {
			if visited.Contains(edge.To) {
				continue
			}

			prev[edge.To] = curr
			visited.Add(edge.To)
			q.Enqueue(edge.To)
		}
	}

	return reconstructPath(prev, source, dest)
}
