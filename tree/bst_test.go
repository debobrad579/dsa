package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func collectInts(traversal func(func(int))) []int {
	result := []int{}
	traversal(func(val int) {
		result = append(result, val)
	})
	return result
}

func setupBST() tree.BinarySearchTree[int] {
	var bst tree.BinarySearchTree[int]
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)

	return bst
	//     4
	//    / \
	//   2   6
	//  / \ / \
	// 1  3 5  7
}

func TestBSTContains(t *testing.T) {
	bst := setupBST()
	assert.True(t, bst.Contains(4))
	assert.True(t, bst.Contains(2))
	assert.True(t, bst.Contains(6))
	assert.True(t, bst.Contains(1))
	assert.True(t, bst.Contains(3))
	assert.True(t, bst.Contains(5))
	assert.True(t, bst.Contains(7))
	assert.False(t, bst.Contains(42))
}

func TestBSTEquals(t *testing.T) {
	bst := setupBST()
	bst2 := setupBST()
	assert.True(t, bst.Equals(&bst2))
	bst2.Insert(10)
	assert.False(t, bst.Equals(&bst2))
}

func TestBSTMinMax(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, 1, bst.Min())
	assert.Equal(t, 7, bst.Max())
}

func TestBSTHeight(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, 3, bst.Height())
}

func TestBSTDelete(t *testing.T) {
	bst := setupBST()
	bst.Delete(6)
	bst.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, collectInts(bst.InOrderTraversal))
}

func TestBSTPreOrderTraversal(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, collectInts(bst.PreOrderTraversal))
}

func TestBSTInOrderTraversal(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(bst.InOrderTraversal))
}

func TestBSTPostOrderTraversal(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, collectInts(bst.PostOrderTraversal))
}

func TestBSTLevelOrderTraversal(t *testing.T) {
	bst := setupBST()
	assert.Equal(t, []int{4, 2, 6, 1, 3, 5, 7}, collectInts(bst.LevelOrderTraversal))
}
