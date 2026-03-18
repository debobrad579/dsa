package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func collectInts(bt *tree.BinaryTree[int], traversal func(func(int))) []int {
	result := []int{}
	traversal(func(val int) {
		result = append(result, val)
	})
	return result
}

func setupTree() tree.BinaryTree[int] {
	var bt tree.BinaryTree[int]
	bt.Insert(4)
	bt.Insert(2)
	bt.Insert(6)
	bt.Insert(1)
	bt.Insert(3)
	bt.Insert(5)
	bt.Insert(7)

	return bt
	//     4
	//    / \
	//   2   6
	//  / \ / \
	// 1  3 5  7
}

func TestPreOrderTraversal(t *testing.T) {
	bt := setupTree()
	result := collectInts(&bt, bt.PreOrderTraversal)
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, result)
}

func TestInOrderTraversal(t *testing.T) {
	bt := setupTree()
	result := collectInts(&bt, bt.InOrderTraversal)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, result)
}

func TestPostOrderTraversal(t *testing.T) {
	bt := setupTree()
	result := collectInts(&bt, bt.PostOrderTraversal)
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, result)
}

func TestTraversalEmptyTree(t *testing.T) {
	var bt tree.BinaryTree[int]
	assert.Equal(t, []int{}, collectInts(&bt, bt.PreOrderTraversal))
	assert.Equal(t, []int{}, collectInts(&bt, bt.InOrderTraversal))
	assert.Equal(t, []int{}, collectInts(&bt, bt.PostOrderTraversal))
}

func TestTraversalSingleNode(t *testing.T) {
	var bt tree.BinaryTree[int]
	bt.Insert(42)
	assert.Equal(t, []int{42}, collectInts(&bt, bt.PreOrderTraversal))
	assert.Equal(t, []int{42}, collectInts(&bt, bt.InOrderTraversal))
	assert.Equal(t, []int{42}, collectInts(&bt, bt.PostOrderTraversal))
}
