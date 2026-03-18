package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func setupAVLTree() tree.AVLTree[int] {
	var avl tree.AVLTree[int]
	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)
	return avl
	//     4
	//    / \
	//   2   6
	//  / \ / \
	// 1  3 5  7
}

func TestAVLContains(t *testing.T) {
	avl := setupTree()
	assert.True(t, avl.Contains(4))
	assert.True(t, avl.Contains(2))
	assert.True(t, avl.Contains(6))
	assert.True(t, avl.Contains(1))
	assert.True(t, avl.Contains(3))
	assert.True(t, avl.Contains(5))
	assert.True(t, avl.Contains(7))
	assert.False(t, avl.Contains(42))
}

func TestAVLEquals(t *testing.T) {
	avl := setupTree()
	avl2 := setupTree()
	assert.True(t, avl.Equals(&avl2))
	avl2.Insert(10)
	assert.False(t, avl.Equals(&avl2))
}

func TestAVLMinMax(t *testing.T) {
	avl := setupTree()
	assert.Equal(t, 1, avl.Min())
	assert.Equal(t, 7, avl.Max())
}

func TestAVLHeight(t *testing.T) {
	avl := setupTree()
	assert.Equal(t, 3, avl.Height())
}

func TestAVLDelete(t *testing.T) {
	avl := setupTree()
	avl.Delete(6)
	avl.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, collectInts(avl.InOrderTraversal))
}
