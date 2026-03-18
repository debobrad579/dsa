package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type AVLTree[T cmp.Ordered] struct {
	root *avlNode[T]
}

type avlNode[T cmp.Ordered] struct {
	val    T
	left   *avlNode[T]
	right  *avlNode[T]
	height int
}

func (n *avlNode[T]) nodeHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *avlNode[T]) updateHeight() {
	n.height = max(n.left.nodeHeight(), n.right.nodeHeight()) + 1
}

func (n *avlNode[T]) balanceFactor() int {
	if n == nil {
		return 0
	}

	return n.left.nodeHeight() - n.right.nodeHeight()
}

func (avl *AVLTree[T]) Height() int {
	return avl.root.nodeHeight()
}

func (avl *AVLTree[T]) Empty() bool {
	return avl.root == nil
}

func (avl *AVLTree[T]) Contains(val T) bool {
	return avl.root.contains(val)
}

func (n *avlNode[T]) contains(val T) bool {
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

func (avl *AVLTree[T]) Equals(other *AVLTree[T]) bool {
	return avl.root.equals(other.root)
}

func (n *avlNode[T]) equals(other *avlNode[T]) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil || n.val != other.val {
		return false
	}

	return n.left.equals(other.left) && n.right.equals(other.right)
}

func (avl *AVLTree[T]) Min() T {
	if avl.root == nil {
		var zero T
		return zero
	}

	return avl.root.min()
}

func (n *avlNode[T]) min() T {
	if n.left == nil {
		return n.val
	}

	return n.left.min()
}

func (avl *AVLTree[T]) Max() T {
	if avl.root == nil {
		var zero T
		return zero
	}

	return avl.root.max()
}

func (n *avlNode[T]) max() T {
	if n.right == nil {
		return n.val
	}

	return n.right.max()
}

func (n *avlNode[T]) leftRotate() *avlNode[T] {
	rightChild := n.right
	n.right = rightChild.left
	rightChild.left = n
	n.updateHeight()
	rightChild.updateHeight()
	return rightChild
}

func (n *avlNode[T]) rightRotate() *avlNode[T] {
	leftChild := n.left
	n.left = leftChild.right
	leftChild.right = n
	n.updateHeight()
	leftChild.updateHeight()
	return leftChild
}

func (n *avlNode[T]) rotate() *avlNode[T] {
	switch n.balanceFactor() {
	case 2:
		if n.left.balanceFactor() >= 0 {
			return n.rightRotate()
		} else {
			n.left = n.left.leftRotate()
			return n.leftRotate()
		}
	case -2:
		if n.right.balanceFactor() <= 0 {
			return n.leftRotate()
		} else {
			n.right = n.right.rightRotate()
			return n.rightRotate()
		}
	default:
		return n
	}
}

func (avl *AVLTree[T]) Insert(val T) {
	avl.root = avl.root.insert(val)
}

func (n *avlNode[T]) insert(val T) *avlNode[T] {
	if n == nil {
		return &avlNode[T]{val: val, height: 1}
	}

	if val < n.val {
		n.left = n.left.insert(val)
	} else if val > n.val {
		n.right = n.right.insert(val)
	}

	n.updateHeight()
	return n.rotate()
}

func (avl *AVLTree[T]) Delete(val T) {
	avl.root = avl.root.delete(val)
}

func (n *avlNode[T]) delete(val T) *avlNode[T] {
	if n == nil {
		return nil
	}

	switch {
	case val < n.val:
		n.left = n.left.delete(val)
	case val > n.val:
		n.right = n.right.delete(val)
	case val == n.val:
		if n.right == nil {
			return n.left
		}

		if n.left == nil {
			return n.right
		}

		if n.left.nodeHeight() > n.right.nodeHeight() {
			successorVal := n.left.max()
			n.val = successorVal
			n.left = n.left.delete(successorVal)
		} else {
			successorVal := n.right.min()
			n.val = successorVal
			n.right = n.right.delete(successorVal)
		}
	}

	n.updateHeight()
	return n.rotate()
}

func (avl *AVLTree[T]) PreOrderTraversal(callback func(val T)) {
	avl.root.preOrderTraversal(callback)
}

func (n *avlNode[T]) preOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	callback(n.val)
	n.left.preOrderTraversal(callback)
	n.right.preOrderTraversal(callback)
}

func (avl *AVLTree[T]) InOrderTraversal(callback func(val T)) {
	avl.root.inOrderTraversal(callback)
}

func (n *avlNode[T]) inOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.inOrderTraversal(callback)
	callback(n.val)
	n.right.inOrderTraversal(callback)
}

func (avl *AVLTree[T]) PostOrderTraversal(callback func(val T)) {
	avl.root.postOrderTraversal(callback)
}

func (n *avlNode[T]) postOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.postOrderTraversal(callback)
	n.right.postOrderTraversal(callback)
	callback(n.val)
}

func (avl *AVLTree[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.Queue[*avlNode[T]]{}
	q.Enqueue(avl.root)

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
