package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func collectInts(bt *tree.BinarySearchTree[int], traversal func(func(int))) []int {
	result := []int{}
	traversal(func(val int) {
		result = append(result, val)
	})
	return result
}

func setupTree() tree.BinarySearchTree[int] {
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

func TestPreOrderTraversal(t *testing.T) {
	bst := setupTree()
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, collectInts(&bst, bst.PreOrderTraversal))
}

func TestInOrderTraversal(t *testing.T) {
	bst := setupTree()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(&bst, bst.InOrderTraversal))
}

func TestPostOrderTraversal(t *testing.T) {
	bst := setupTree()
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, collectInts(&bst, bst.PostOrderTraversal))
}

func TestLevelOrderTraversal(t *testing.T) {
	bst := setupTree()
	assert.Equal(t, []int{4, 2, 6, 1, 3, 5, 7}, collectInts(&bst, bst.LevelOrderTraversal))
}

func TestTraversalEmptyTree(t *testing.T) {
	var bst tree.BinarySearchTree[int]
	assert.Equal(t, []int{}, collectInts(&bst, bst.PreOrderTraversal))
	assert.Equal(t, []int{}, collectInts(&bst, bst.InOrderTraversal))
	assert.Equal(t, []int{}, collectInts(&bst, bst.PostOrderTraversal))
	assert.Equal(t, []int{}, collectInts(&bst, bst.LevelOrderTraversal))
}

func TestTraversalSingleNode(t *testing.T) {
	var bst tree.BinarySearchTree[int]
	bst.Insert(42)
	assert.Equal(t, []int{42}, collectInts(&bst, bst.PreOrderTraversal))
	assert.Equal(t, []int{42}, collectInts(&bst, bst.InOrderTraversal))
	assert.Equal(t, []int{42}, collectInts(&bst, bst.PostOrderTraversal))
	assert.Equal(t, []int{42}, collectInts(&bst, bst.LevelOrderTraversal))
}

func TestDelete(t *testing.T) {
	bst := setupTree()
	bst.Delete(6)
	bst.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, collectInts(&bst, bst.InOrderTraversal))
}
