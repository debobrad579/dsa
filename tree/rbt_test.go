package tree_test

import (
	"testing"

	"github.com/debobrad579/dsa/tree"
	"github.com/stretchr/testify/assert"
)

func setupRBT() tree.RedBlackTree[int] {
	var rbt tree.RedBlackTree[int]
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

func TestRBTEquals(t *testing.T) {
	rbt := setupRBT()
	rbt2 := setupRBT()
	assert.True(t, rbt.Equals(&rbt2))
	rbt2.Insert(10)
	assert.False(t, rbt.Equals(&rbt2))
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
}

func TestRBTBalancingOnInsert(t *testing.T) {
	rbt := setupRBT()
	assert.Equal(t, 4, rbt.Height())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, collectInts(rbt.InOrderTraversal))
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
