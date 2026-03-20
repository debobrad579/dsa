package path

import (
	"slices"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/queue"
	"github.com/debobrad579/dsa/set"
)

func BreadthFirstSearchList(graph graph.AdjacencyList, source, target int) []int {
	q := queue.New[int]()
	q.Enqueue(source)
	visited := set.New(source)
	prev := make([]int, len(graph))
	for i := range len(graph) {
		prev[i] = -1
	}

	for !q.Empty() {
		curr := q.Deque()
		if curr == target {
			break
		}

		for _, edge := range graph[curr] {
			if visited.Contains(edge.To) {
				continue
			}

			prev[edge.To] = curr
			visited.Add(edge.To)
			q.Enqueue(edge.To)
		}
	}

	path := make([]int, 0, len(graph))
	for curr := target; prev[curr] != -1; {
		path = append(path, curr)
		curr = prev[curr]
	}

	if len(path) != 0 {
		path = append(path, source)
		slices.Reverse(path)
	}

	return path
}
