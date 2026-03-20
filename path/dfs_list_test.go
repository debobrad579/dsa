package path_test

import (
	"testing"

	"github.com/debobrad579/dsa/graph"
	"github.com/debobrad579/dsa/path"
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

func TestDFSListSourceEqualsTarget(t *testing.T) {
	result := path.DepthFirstSearchList(linearList, 0, 0)
	assert.Equal(t, []int{0}, result)
}

func TestDFSListNoPath(t *testing.T) {
	result := path.DepthFirstSearchList(linearList, 3, 0)
	assert.Equal(t, []int{}, result)
}

func TestDFSListDirectNeighbour(t *testing.T) {
	result := path.DepthFirstSearchList(linearList, 0, 1)
	assert.Equal(t, []int{0, 1}, result)
}

func TestDFSListLinearPath(t *testing.T) {
	result := path.DepthFirstSearchList(linearList, 0, 3)
	assert.Equal(t, []int{0, 1, 2, 3}, result)
}

func TestDFSListShortcutPath(t *testing.T) {
	result := path.DepthFirstSearchList(shortcutList, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestDFSListDeadEnd(t *testing.T) {
	result := path.DepthFirstSearchList(shortcutList, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestDFSListDiamond(t *testing.T) {
	result := path.DepthFirstSearchList(diamondList, 0, 3)
	via1 := []int{0, 1, 3}
	via2 := []int{0, 2, 3}
	assert.True(t,
		assert.ObjectsAreEqual(via1, result) || assert.ObjectsAreEqual(via2, result),
		"expected path via 1 or via 2, got %v", result,
	)
}

func TestDFSListDetour(t *testing.T) {
	result := path.DepthFirstSearchList(detourList, 0, 3)
	valid := [][]int{
		{0, 2, 3},
		{0, 1, 4, 5, 3},
	}
	assert.Contains(t, valid, result)
}
