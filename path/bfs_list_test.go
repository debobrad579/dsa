package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/debobrad579/dsa/path"
)

func TestBFSListShortcutPath(t *testing.T) {
	result := path.BreadthFirstSearchList(shortcutList, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestBFSListShortcutDeadEnd(t *testing.T) {
	result := path.BreadthFirstSearchList(shortcutList, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestBFSListShortcutNoPath(t *testing.T) {
	result := path.BreadthFirstSearchList(shortcutList, 2, 3)
	assert.Equal(t, []int{}, result)
}

func TestBFSDiamondShortestPath(t *testing.T) {
	result := path.BreadthFirstSearchList(diamondList, 0, 3)
	via1 := []int{0, 1, 3}
	via2 := []int{0, 2, 3}
	assert.True(t,
		assert.ObjectsAreEqual(via1, result) || assert.ObjectsAreEqual(via2, result),
		"expected path via 1 or via 2, got %v", result,
	)
}

func TestBFSDiamondDirectHop(t *testing.T) {
	result := path.BreadthFirstSearchList(diamondList, 1, 3)
	assert.Equal(t, []int{1, 3}, result)
}

func TestBFSListNoPath(t *testing.T) {
	result := path.BreadthFirstSearchList(diamondList, 3, 0)
	assert.Equal(t, []int{}, result)
}

func TestBFSListDetourShortestPath(t *testing.T) {
	result := path.BreadthFirstSearchList(detourList, 0, 3)
	assert.Equal(t, []int{0, 2, 3}, result)
}

func TestBFSListDetourLongPathExists(t *testing.T) {
	result := path.BreadthFirstSearchList(detourList, 1, 3)
	assert.Equal(t, []int{1, 4, 5, 3}, result)
}

func TestBFSListDetourNoPath(t *testing.T) {
	result := path.BreadthFirstSearchList(detourList, 3, 0)
	assert.Equal(t, []int{}, result)
}
