package path_test

import (
	"testing"

	"github.com/debobrad579/dsa/path"
	"github.com/stretchr/testify/assert"
)

func TestDFSMatrixShortcutPath(t *testing.T) {
	result := path.DepthFirstSearchMatrix(shortcut, 0, 3)
	assert.Equal(t, []int{0, 1, 3}, result)
}

func TestDFSMatrixDeadEnd(t *testing.T) {
	result := path.DepthFirstSearchMatrix(shortcut, 0, 2)
	assert.Equal(t, []int{0, 2}, result)
}

func TestDFSMatrixDiamond(t *testing.T) {
	result := path.DepthFirstSearchMatrix(diamond, 0, 3)
	via1 := []int{0, 1, 3}
	via2 := []int{0, 2, 3}
	assert.True(t,
		assert.ObjectsAreEqual(via1, result) || assert.ObjectsAreEqual(via2, result),
		"expected path via 1 or via 2, got %v", result,
	)
}

func TestDFSMatrixDetour(t *testing.T) {
	result := path.DepthFirstSearchMatrix(detour, 0, 3)
	valid := [][]int{
		{0, 2, 3},
		{0, 1, 4, 5, 3},
	}
	assert.Contains(t, valid, result)
}
