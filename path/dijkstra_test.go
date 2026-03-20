package path_test

import (
	"testing"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/path"
	"github.com/stretchr/testify/assert"
)

var weightedDiamondList = graph.AdjacencyList{
	{{To: 1, Weight: 1}, {To: 2, Weight: 1}},
	{{To: 3, Weight: 10}},
	{{To: 3, Weight: 1}},
	{},
}

var weightedDetourList = graph.AdjacencyList{
	{{To: 1, Weight: 1}, {To: 2, Weight: 1}},
	{{To: 4, Weight: 1}},
	{{To: 3, Weight: 1}},
	{},
	{{To: 5, Weight: 1}},
	{{To: 3, Weight: 1}},
}

var weightedShortcutList = graph.AdjacencyList{
	{{To: 1, Weight: 1}, {To: 2, Weight: 1}},
	{{To: 3, Weight: 1}},
	{},
	{},
}

func TestDijkstraShortcutPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedShortcutList, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestDijkstraShortcutDeadEnd(t *testing.T) {
	result := path.DijkstraShortestPath(weightedShortcutList, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestDijkstraShortcutNoPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedShortcutList, 2, 3)
	assert.Equal(t, []int{}, result)
}

func TestDijkstraDiamondShortestPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDiamondList, 0, 3)
	assert.Equal(t, []int{0, 2, 3}, result)
}

func TestDijkstraDiamondDirectHop(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDiamondList, 1, 3)
	assert.Equal(t, []int{1, 3}, result)
}

func TestDijkstraDiamondNoPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDiamondList, 3, 0)
	assert.Equal(t, []int{}, result)
}

func TestDijkstraDetourShortestPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDetourList, 0, 3)
	assert.Equal(t, []int{0, 2, 3}, result)
}

func TestDijkstraDetourLongPathExists(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDetourList, 1, 3)
	assert.Equal(t, []int{1, 4, 5, 3}, result)
}

func TestDijkstraDetourNoPath(t *testing.T) {
	result := path.DijkstraShortestPath(weightedDetourList, 3, 0)
	assert.Equal(t, []int{}, result)
}
