package path_test

import (
	"slices"
	"testing"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/path"
	"github.com/stretchr/testify/assert"
)

var negativeWeightsList = graph.AdjacencyList{
	{{To: 1, Weight: 1}, {To: 2, Weight: 4}},
	{{To: 2, Weight: -3}, {To: 3, Weight: 2}},
	{{To: 3, Weight: 3}},
	{},
}

var negativeCycleList = graph.AdjacencyList{
	{{To: 1, Weight: 1}},
	{{To: 2, Weight: -1}},
	{{To: 1, Weight: -1}},
}

var unreachableNegativeCycleList = graph.AdjacencyList{
	{{To: 1, Weight: 1}},
	{},
	{{To: 3, Weight: -1}},
	{{To: 2, Weight: -1}},
}

var multiplePathsList = graph.AdjacencyList{
	{{To: 1, Weight: 1}, {To: 2, Weight: 1}},
	{{To: 3, Weight: 1}},
	{{To: 3, Weight: 1}},
	{},
}

func TestBellmanFord_SimpleShortestPath(t *testing.T) {
	result, err := path.BellmanFord(weightedDiamondList, 0, 3)
	assert.NoError(t, err)
	assert.Equal(t, []int{0, 2, 3}, result)
}

func TestBellmanFord_NegativeWeights(t *testing.T) {
	result, err := path.BellmanFord(negativeWeightsList, 0, 3)
	assert.NoError(t, err)
	assert.Equal(t, []int{0, 1, 2, 3}, result)
}

func TestBellmanFord_NegativeCycle(t *testing.T) {
	result, err := path.BellmanFord(negativeCycleList, 0, 2)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestBellmanFord_UnreachableNegativeCycle(t *testing.T) {
	result, err := path.BellmanFord(unreachableNegativeCycleList, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, []int{0, 1}, result)
}

func TestBellmanFord_NoPath(t *testing.T) {
	result, err := path.BellmanFord(weightedShortcutList, 2, 3)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestBellmanFord_SameSourceDestination(t *testing.T) {
	result, err := path.BellmanFord(weightedDiamondList, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, []int{0}, result)
}

func TestBellmanFord_MultipleValidPaths(t *testing.T) {
	result, err := path.BellmanFord(multiplePathsList, 0, 3)
	assert.NoError(t, err)

	valid1 := []int{0, 1, 3}
	valid2 := []int{0, 2, 3}

	assert.True(t,
		slices.Equal(result, valid1) || slices.Equal(result, valid2),
	)
}
