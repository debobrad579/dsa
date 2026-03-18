package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type BinarySearchTree[T cmp.Ordered] struct {
	root *bstNode[T]
}

type bstNode[T cmp.Ordered] struct {
	val   T
	left  *bstNode[T]
	right *bstNode[T]
}

func (bst *BinarySearchTree[T]) Empty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree[T]) Contains(val T) bool {
	return bst.root.contains(val)
}

func (n *bstNode[T]) contains(val T) bool {
	if n == nil {
		return false
	}

	if n.val == val {
		return true
	}

	if val < n.val {
		return n.left.contains(val)
	}

	return n.right.contains(val)
}

func (bst *BinarySearchTree[T]) Equals(other *BinarySearchTree[T]) bool {
	return bst.root.equals(other.root)
}

func (n *bstNode[T]) equals(other *bstNode[T]) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil || n.val != other.val {
		return false
	}

	return n.left.equals(other.left) && n.right.equals(other.right)
}

func (bst *BinarySearchTree[T]) Min() T {
	if bst.root == nil {
		var zero T
		return zero
	}

	return bst.root.min()
}

func (n *bstNode[T]) min() T {
	if n.left == nil {
		return n.val
	}

	return n.left.min()
}

func (bst *BinarySearchTree[T]) Max() T {
	if bst.root == nil {
		var zero T
		return zero
	}

	return bst.root.max()
}

func (n *bstNode[T]) max() T {
	if n.right == nil {
		return n.val
	}

	return n.right.max()
}

func (bst *BinarySearchTree[T]) Height() int {
	return bst.root.height()
}

func (n *bstNode[T]) height() int {
	if n == nil {
		return 0
	}

	return max(n.left.height(), n.right.height()) + 1
}

func (bst *BinarySearchTree[T]) Insert(val T) {
	bst.root = bst.root.insert(val)
}

func (n *bstNode[T]) insert(val T) *bstNode[T] {
	if n == nil {
		return &bstNode[T]{val: val}
	}

	if val < n.val {
		n.left = n.left.insert(val)
	} else if val > n.val {
		n.right = n.right.insert(val)
	}
	return n
}

func (bst *BinarySearchTree[T]) Delete(val T) {
	bst.root = bst.root.delete(val)
}

func (n *bstNode[T]) delete(val T) *bstNode[T] {
	if n == nil {
		return nil
	}

	if val < n.val {
		n.left = n.left.delete(val)
		return n
	}

	if val > n.val {
		n.right = n.right.delete(val)
		return n
	}

	if n.right == nil {
		return n.left
	}

	if n.left == nil {
		return n.right
	}

	successorVal := n.right.min()
	n.val = successorVal
	n.right = n.right.delete(successorVal)
	return n
}

func (bst *BinarySearchTree[T]) PreOrderTraversal(callback func(val T)) {
	bst.root.preOrderTraversal(callback)
}

func (n *bstNode[T]) preOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	callback(n.val)
	n.left.preOrderTraversal(callback)
	n.right.preOrderTraversal(callback)
}

func (bst *BinarySearchTree[T]) InOrderTraversal(callback func(val T)) {
	bst.root.inOrderTraversal(callback)
}

func (n *bstNode[T]) inOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.inOrderTraversal(callback)
	callback(n.val)
	n.right.inOrderTraversal(callback)
}

func (bst *BinarySearchTree[T]) PostOrderTraversal(callback func(val T)) {
	bst.root.postOrderTraversal(callback)
}

func (n *bstNode[T]) postOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.postOrderTraversal(callback)
	n.right.postOrderTraversal(callback)
	callback(n.val)
}

func (bst *BinarySearchTree[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.Queue[*bstNode[T]]{}
	q.Enqueue(bst.root)

	for !q.Empty() {
		n := q.Deque()
		if n == nil {
			continue
		}

		if n.left != nil {
			q.Enqueue(n.left)
		}

		if n.right != nil {
			q.Enqueue(n.right)
		}

		callback(n.val)
	}
}
