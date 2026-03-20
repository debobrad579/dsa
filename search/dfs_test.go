package search_test

import (
	"testing"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/search"
	"github.com/stretchr/testify/assert"
)

var linearList = graph.AdjacencyList{
	{{To: 1}},
	{{To: 2}},
	{{To: 3}},
	{},
}

var shortcutList = graph.AdjacencyList{
	{{To: 1}, {To: 2}},
	{{To: 3}},
	{},
	{},
}

var diamondList = graph.AdjacencyList{
	{{To: 1}, {To: 2}},
	{{To: 3}},
	{{To: 3}},
	{},
}

var detourList = graph.AdjacencyList{
	{{To: 1}, {To: 2}},
	{{To: 4}},
	{{To: 3}},
	{},
	{{To: 5}},
	{{To: 3}},
}

func TestDFSSourceEqualsTarget(t *testing.T) {
	result := search.DepthFirstSearch(linearList, 0, 0)
	assert.Equal(t, []int{0}, result)
}

func TestDFSNoPath(t *testing.T) {
	result := search.DepthFirstSearch(linearList, 3, 0)
	assert.Equal(t, []int{}, result)
}

func TestDFSDirectNeighbour(t *testing.T) {
	result := search.DepthFirstSearch(linearList, 0, 1)
	assert.Equal(t, []int{0, 1}, result)
}

func TestDFSLinearPath(t *testing.T) {
	result := search.DepthFirstSearch(linearList, 0, 3)
	assert.Equal(t, []int{0, 1, 2, 3}, result)
}

func TestDFSShortcutPath(t *testing.T) {
	result := search.DepthFirstSearch(shortcutList, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestDFSDeadEnd(t *testing.T) {
	result := search.DepthFirstSearch(shortcutList, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestDFSDiamond(t *testing.T) {
	result := search.DepthFirstSearch(diamondList, 0, 3)
	via1 := []int{0, 1, 3}
	via2 := []int{0, 2, 3}
	assert.True(t,
		assert.ObjectsAreEqual(via1, result) || assert.ObjectsAreEqual(via2, result),
		"expected path via 1 or via 2, got %v", result,
	)
}

func TestDFSDetour(t *testing.T) {
	result := search.DepthFirstSearch(detourList, 0, 3)
	valid := [][]int{
		{0, 2, 3},
		{0, 1, 4, 5, 3},
	}
	assert.Contains(t, valid, result)
}
