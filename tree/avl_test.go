package tree_test

import (
	"slices"
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func setupAVLTree() tree.BinarySearchTree[int] {
	avl := tree.NewAVLTree[int]()
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

func TestAVLDelete(t *testing.T) {
	avl := setupAVLTree()
	avl.Delete(6)
	avl.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, slices.Collect(avl.InOrderTraversal()))
}

func TestAVLBalancingOnInsert(t *testing.T) {
	avl := setupAVLTree()
	assert.Equal(t, 3, avl.Height())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, slices.Collect(avl.InOrderTraversal()))
}

func TestAVLBalancingOnDelete(t *testing.T) {
	avl := setupAVLTree()
	avl.Delete(1)
	avl.Delete(2)
	avl.Delete(3)
	avl.Delete(4)
	assert.Equal(t, 2, avl.Height())
	assert.Equal(t, []int{5, 6, 7}, slices.Collect(avl.InOrderTraversal()))
}
