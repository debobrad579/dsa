package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type RedBlackTree[T cmp.Ordered] struct {
	root *rbtNode[T]
}

type rbtNode[T cmp.Ordered] struct {
	val    T
	left   *rbtNode[T]
	right  *rbtNode[T]
	parent *rbtNode[T]
	red    bool
}

func (n *rbtNode[T]) isRed() bool {
	if n == nil {
		return false
	}

	return n.red
}

func (rbt *RedBlackTree[T]) Empty() bool {
	return rbt.root == nil
}

func (rbt *RedBlackTree[T]) Contains(val T) bool {
	return rbt.root.contains(val)
}

func (n *rbtNode[T]) contains(val T) bool {
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

func (rbt *RedBlackTree[T]) Equals(other *RedBlackTree[T]) bool {
	return rbt.root.equals(other.root)
}

func (n *rbtNode[T]) equals(other *rbtNode[T]) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil || n.val != other.val {
		return false
	}

	return n.left.equals(other.left) && n.right.equals(other.right)
}

func (rbt *RedBlackTree[T]) Min() T {
	if rbt.root == nil {
		var zero T
		return zero
	}

	return rbt.root.min()
}

func (n *rbtNode[T]) min() T {
	if n.left == nil {
		return n.val
	}

	return n.left.min()
}

func (rbt *RedBlackTree[T]) Max() T {
	if rbt.root == nil {
		var zero T
		return zero
	}

	return rbt.root.max()
}

func (n *rbtNode[T]) max() T {
	if n.right == nil {
		return n.val
	}

	return n.right.max()
}

func (rbt *RedBlackTree[T]) Height() int {
	return rbt.root.height()
}

func (n *rbtNode[T]) height() int {
	if n == nil {
		return 0
	}

	return max(n.left.height(), n.right.height()) + 1
}

func (rbt *RedBlackTree[T]) rotateLeft(pivotParent *rbtNode[T]) {
	if pivotParent == nil || pivotParent.right == nil {
		return
	}

	pivot := pivotParent.right
	pivotParent.right = pivot.left
	if pivot.left != nil {
		pivot.left.parent = pivotParent
	}
	pivot.parent = pivotParent.parent

	switch pivotParent {
	case rbt.root:
		rbt.root = pivot
	case pivotParent.parent.left:
		pivotParent.parent.left = pivot
	case pivotParent.parent.right:
		pivotParent.parent.right = pivot
	}

	pivot.left = pivotParent
	pivotParent.parent = pivot
}

func (rbt *RedBlackTree[T]) rotateRight(pivotParent *rbtNode[T]) {
	if pivotParent == nil || pivotParent.left == nil {
		return
	}

	pivot := pivotParent.left
	pivotParent.left = pivot.right
	if pivot.right != nil {
		pivot.right.parent = pivotParent
	}
	pivot.parent = pivotParent.parent

	switch pivotParent {
	case rbt.root:
		rbt.root = pivot
	case pivotParent.parent.left:
		pivotParent.parent.left = pivot
	case pivotParent.parent.right:
		pivotParent.parent.right = pivot
	}

	pivot.right = pivotParent
	pivotParent.parent = pivot
}

func (rbt *RedBlackTree[T]) Insert(val T) {
	newNode := &rbtNode[T]{val: val, red: true}

	var parent *rbtNode[T] = nil
	for currentNode := rbt.root; currentNode != nil; {
		if newNode.val == currentNode.val {
			return
		}

		parent = currentNode
		if newNode.val < currentNode.val {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	newNode.parent = parent
	if newNode.parent == nil {
		rbt.root = newNode
	} else {
		if newNode.val < parent.val {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}

	rbt.fixInsert(newNode)
	rbt.root.red = false
}

func (rbt *RedBlackTree[T]) fixInsert(newNode *rbtNode[T]) {
	for newNode != rbt.root && newNode.parent.red {
		parent := newNode.parent
		grandparent := parent.parent

		switch parent {
		case grandparent.right:
			uncle := grandparent.left
			if uncle.isRed() {
				uncle.red = false
				parent.red = false
				grandparent.red = true
				newNode = grandparent
			} else {
				if newNode == parent.left {
					newNode = parent
					rbt.rotateRight(newNode)
					parent = newNode.parent
				}
				parent.red = false
				grandparent.red = true
				rbt.rotateLeft(grandparent)
			}
		case grandparent.left:
			uncle := grandparent.right
			if uncle.isRed() {
				uncle.red = false
				parent.red = false
				grandparent.red = true
				newNode = grandparent
			} else {
				if newNode == parent.right {
					newNode = parent
					rbt.rotateLeft(newNode)
					parent = newNode.parent
				}
				parent.red = false
				grandparent.red = true
				rbt.rotateRight(grandparent)
			}
		}
	}
}

func (rbt *RedBlackTree[T]) Delete(val T) {
	rbt.root = rbt.root.delete(val)
}

func (n *rbtNode[T]) delete(val T) *rbtNode[T] {
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

func (rbt *RedBlackTree[T]) PreOrderTraversal(callback func(val T)) {
	rbt.root.preOrderTraversal(callback)
}

func (n *rbtNode[T]) preOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	callback(n.val)
	n.left.preOrderTraversal(callback)
	n.right.preOrderTraversal(callback)
}

func (rbt *RedBlackTree[T]) InOrderTraversal(callback func(val T)) {
	rbt.root.inOrderTraversal(callback)
}

func (n *rbtNode[T]) inOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.inOrderTraversal(callback)
	callback(n.val)
	n.right.inOrderTraversal(callback)
}

func (rbt *RedBlackTree[T]) PostOrderTraversal(callback func(val T)) {
	rbt.root.postOrderTraversal(callback)
}

func (n *rbtNode[T]) postOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.postOrderTraversal(callback)
	n.right.postOrderTraversal(callback)
	callback(n.val)
}

func (rbt *RedBlackTree[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.Queue[*rbtNode[T]]{}
	q.Enqueue(rbt.root)

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
