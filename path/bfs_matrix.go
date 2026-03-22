package path

import (
	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/queue"
	"github.com/debobrad579/dsa/set"
)

func BreadthFirstSearchMatrix(g graph.AdjacencyMatrix, source, dest int) []int {
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

		for i, edge := range g[curr] {
			if edge == 0 || visited.Contains(i) {
				continue
			}

			prev[i] = curr
			visited.Add(i)
			q.Enqueue(i)
		}
	}

	return reconstructPath(prev, source, dest)
}
