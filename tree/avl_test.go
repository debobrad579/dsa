package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func setupAVLTree() tree.AVLTree[int] {
	var avl tree.AVLTree[int]
	for i := 1; i <= 7; i++ {
		avl.Insert(i)
	}
	return avl
	//     4
	//    / \
	//   2   6
	//  / \ / \
	// 1  3 5  7
}

func TestAVLContains(t *testing.T) {
	avl := setupAVLTree()
	assert.True(t, avl.Contains(1))
	assert.True(t, avl.Contains(2))
	assert.True(t, avl.Contains(3))
	assert.True(t, avl.Contains(4))
	assert.True(t, avl.Contains(5))
	assert.True(t, avl.Contains(6))
	assert.True(t, avl.Contains(7))
	assert.False(t, avl.Contains(42))
}

func TestAVLEquals(t *testing.T) {
	avl := setupAVLTree()
	avl2 := setupAVLTree()
	assert.True(t, avl.Equals(&avl2))
	avl2.Insert(10)
	assert.False(t, avl.Equals(&avl2))
}

func TestAVLMinMax(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, 1, avl.Min())
	assert.Equal(t, 7, avl.Max())
}

func TestAVLDelete(t *testing.T) {
	avl := setupAVLTree()
	avl.Delete(6)
	avl.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, collectInts(avl.InOrderTraversal))
}

func TestAVLBalancingOnInsert(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, 3, avl.Height())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(avl.InOrderTraversal))
}

func TestAVLBalancingOnDelete(t *testing.T) {
	avl := setupAVLTree()
	avl.Delete(1)
	avl.Delete(2)
	avl.Delete(3)
	avl.Delete(4)
	assert.Equal(t, 2, avl.Height())
	assert.Equal(t, []int{5, 6, 7}, collectInts(avl.InOrderTraversal))
}

func TestAVLPreOrderTraversal(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, collectInts(avl.PreOrderTraversal))
}

func TestAVLInOrderTraversal(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(avl.InOrderTraversal))
}

func TestAVLPostOrderTraversal(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, collectInts(avl.PostOrderTraversal))
}

func TestAVLLevelOrderTraversal(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, []int{4, 2, 6, 1, 3, 5, 7}, collectInts(avl.LevelOrderTraversal))
}
