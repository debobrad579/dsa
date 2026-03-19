package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func collectInts(traversal func(func(int))) []int {
	result := []int{}
	traversal(func(val int) {
		result = append(result, val)
	})
	return result
}

func assertValidRBT(t *testing.T, rbt *redBlackTree[int]) {
	assert.False(t, rbt.root.isRed())

	var validate func(n *rbtNode[int]) int
	validate = func(n *rbtNode[int]) int {
		if n == nil {
			return 1
		}

		if n.isRed() {
			assert.False(t, n.left.isRed(), "double red at left child of %v", n.val)
			assert.False(t, n.right.isRed(), "double red at right child of %v", n.val)
		}

		leftBH := validate(n.left)
		rightBH := validate(n.right)
		assert.Equal(t, leftBH, rightBH, "black-height mismatch at %v", n.val)

		if n.isRed() {
			return leftBH
		}

		return leftBH + 1
	}

	validate(rbt.root)
}

func setupRBT() redBlackTree[int] {
	var rbt redBlackTree[int]
	for i := 1; i <= 7; i++ {
		rbt.Insert(i)
	}
	return rbt
	//   2
	//  / \
	// 1   4
	//    / \
	//   3   6
	//      / \
	//     5   7
}

func TestRBTContains(t *testing.T) {
	rbt := setupRBT()
	assert.True(t, rbt.Contains(1))
	assert.True(t, rbt.Contains(2))
	assert.True(t, rbt.Contains(3))
	assert.True(t, rbt.Contains(4))
	assert.True(t, rbt.Contains(5))
	assert.True(t, rbt.Contains(6))
	assert.True(t, rbt.Contains(7))
	assert.False(t, rbt.Contains(42))
}

func TestRBTMinMax(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, 1, rbt.Min())
	assert.Equal(t, 7, rbt.Max())
}

func TestRBTDelete(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(6)
	rbt.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, collectInts(rbt.InOrderTraversal))
	assertValidRBT(t, &rbt)
}

func TestRBTDeleteRoot(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(2)
	assertValidRBT(t, &rbt)
	assert.Equal(t, []int{1, 3, 4, 5, 6, 7}, collectInts(rbt.InOrderTraversal))
}

func TestRBTDeleteRootSingleNode(t *testing.T) {
	var rbt redBlackTree[int]
	rbt.Insert(1)
	rbt.Delete(1)
	assert.True(t, rbt.Empty())
}

func TestRBTBalancingOnInsert(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, 4, rbt.Height())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(rbt.InOrderTraversal))
	assertValidRBT(t, &rbt)
}

func TestRBTBalancingOnDelete(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(1)
	assert.Equal(t, 3, rbt.Height())
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, collectInts(rbt.InOrderTraversal))
	assertValidRBT(t, &rbt)
}

func TestRBTPreOrderTraversal(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, []int{2, 1, 4, 3, 6, 5, 7}, collectInts(rbt.PreOrderTraversal))
}

func TestRBTInOrderTraversal(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(rbt.InOrderTraversal))
}

func TestRBTPostOrderTraversal(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, []int{1, 3, 5, 7, 6, 4, 2}, collectInts(rbt.PostOrderTraversal))
}

func TestRBTLevelOrderTraversal(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, []int{2, 1, 4, 3, 6, 5, 7}, collectInts(rbt.LevelOrderTraversal))
}
