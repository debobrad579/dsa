package tree

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestRBTDelete(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(6)
	rbt.Delete(4)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, slices.Collect(rbt.InOrderTraversal()))
	assertValidRBT(t, &rbt)
}

func TestRBTDeleteRoot(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(2)
	assertValidRBT(t, &rbt)
	assert.Equal(t, []int{1, 3, 4, 5, 6, 7}, slices.Collect(rbt.InOrderTraversal()))
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
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, slices.Collect(rbt.InOrderTraversal()))
	assertValidRBT(t, &rbt)
}

func TestRBTBalancingOnDelete(t *testing.T) {
	rbt := setupRBT()
	rbt.Delete(1)
	assert.Equal(t, 3, rbt.Height())
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, slices.Collect(rbt.InOrderTraversal()))
	assertValidRBT(t, &rbt)
}
