package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/path"
)

var shortcut = [][]int{
	{0, 1, 1, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var diamond = [][]int{
	{0, 1, 1, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 0, 0},
}

var detour = [][]int{
	{0, 1, 1, 0, 0, 0},
	{0, 0, 0, 0, 1, 0},
	{0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0},
}

func TestShortcutPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(shortcut, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestShortcutDeadEnd(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(shortcut, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestShortcutNoPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(shortcut, 2, 3)
	assert.Equal(t, []int{}, result)
}

func TestDiamondShortestPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(diamond, 0, 3)
	via1 := []int{0, 1, 3}
	via2 := []int{0, 2, 3}
	assert.True(t,
		assert.ObjectsAreEqual(via1, result) || assert.ObjectsAreEqual(via2, result),
		"expected path via 1 or via 2, got %v", result,
	)
}

func TestDiamondDirectHop(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(diamond, 1, 3)
	assert.Equal(t, []int{1, 3}, result)
}

func TestDiamondNoPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(diamond, 3, 0)
	assert.Equal(t, []int{}, result)
}

func TestDetourShortestPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(detour, 0, 3)
	assert.Equal(t, []int{0, 2, 3}, result)
}

func TestDetourLongPathExists(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(detour, 1, 3)
	assert.Equal(t, []int{1, 4, 5, 3}, result)
}

func TestDetourNoPath(t *testing.T) {
	result := path.BreadthFirstSearchMatrix(detour, 3, 0)
	assert.Equal(t, []int{}, result)
}
