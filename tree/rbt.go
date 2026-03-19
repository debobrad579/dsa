package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

func NewRedBlackTree[T cmp.Ordered]() BinarySearchTree[T] {
	return &redBlackTree[T]{}
}

type redBlackTree[T cmp.Ordered] struct {
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

func (rbt *redBlackTree[T]) Empty() bool {
	return rbt.root == nil
}

func (rbt *redBlackTree[T]) Contains(val T) bool {
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

func (rbt *redBlackTree[T]) Equals(other BinarySearchTree[T]) bool {
	o, ok := other.(*redBlackTree[T])
	if !ok {
		return false
	}

	return rbt.root.equals(o.root)
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

func (rbt *redBlackTree[T]) Min() T {
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

func (rbt *redBlackTree[T]) Max() T {
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

func (rbt *redBlackTree[T]) Height() int {
	return rbt.root.height()
}

func (n *rbtNode[T]) height() int {
	if n == nil {
		return 0
	}

	return max(n.left.height(), n.right.height()) + 1
}

func (rbt *redBlackTree[T]) rotateLeft(pivotParent *rbtNode[T]) {
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

func (rbt *redBlackTree[T]) rotateRight(pivotParent *rbtNode[T]) {
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

func (rbt *redBlackTree[T]) Insert(val T) {
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

func (rbt *redBlackTree[T]) fixInsert(n *rbtNode[T]) {
	for n != rbt.root && n.parent.red {
		parent := n.parent
		grandparent := parent.parent

		switch parent {
		case grandparent.right:
			uncle := grandparent.left

			if uncle.isRed() {
				uncle.red = false
				parent.red = false
				grandparent.red = true
				n = grandparent
			} else {
				if n == parent.left {
					n = parent
					rbt.rotateRight(n)
					parent = n.parent
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
				n = grandparent
			} else {
				if n == parent.right {
					n = parent
					rbt.rotateLeft(n)
					parent = n.parent
				}

				parent.red = false
				grandparent.red = true
				rbt.rotateRight(grandparent)
			}
		}
	}
}

func (rbt *redBlackTree[T]) Delete(val T) {
	n := rbt.root

	for n != nil {
		if val < n.val {
			n = n.left
		} else if val > n.val {
			n = n.right
		} else {
			break
		}
	}

	if n == nil {
		return
	}

	if n.left != nil && n.right != nil {
		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}

		n.val = successor.val
		n = successor
	}

	if n.parent == nil {
		child := n.right
		if child != nil {
			child.parent = nil
		}
		rbt.root = child
		return
	}

	child := n.left
	if child == nil {
		child = n.right
	}

	if child == nil {
		if !n.isRed() {
			rbt.fixDelete(n)
		}

		if n == n.parent.left {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}

		n.parent = nil
		return
	}

	child.parent = n.parent

	if n == n.parent.left {
		n.parent.left = child
	} else {
		n.parent.right = child
	}

	if !n.isRed() {
		rbt.fixDelete(child)
	}
}

func (rbt *redBlackTree[T]) fixDelete(n *rbtNode[T]) {
	for n != rbt.root && !n.isRed() {
		parent := n.parent

		switch n {
		case parent.left:
			sibling := parent.right

			if sibling.isRed() {
				sibling.red = false
				parent.red = true
				rbt.rotateLeft(parent)
				sibling = parent.right
			}

			if !sibling.left.isRed() && !sibling.right.isRed() {
				sibling.red = true
				n = parent
				parent = n.parent
			} else {
				if !sibling.right.isRed() {
					sibling.left.red = false
					sibling.red = true
					rbt.rotateRight(sibling)
					sibling = parent.right
				}

				sibling.red = parent.red
				parent.red = false
				sibling.right.red = false
				rbt.rotateLeft(parent)
				n = rbt.root
			}
		case parent.right:
			sibling := parent.left

			if sibling.isRed() {
				sibling.red = false
				parent.red = true
				rbt.rotateRight(parent)
				sibling = parent.left
			}

			if !sibling.right.isRed() && !sibling.left.isRed() {
				sibling.red = true
				n = parent
				parent = n.parent
			} else {
				if !sibling.left.isRed() {
					sibling.right.red = false
					sibling.red = true
					rbt.rotateLeft(sibling)
					sibling = parent.left
				}

				sibling.red = parent.red
				parent.red = false
				sibling.left.red = false
				rbt.rotateRight(parent)
				n = rbt.root
			}
		}
	}

	n.red = false
}

func (rbt *redBlackTree[T]) PreOrderTraversal(callback func(val T)) {
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

func (rbt *redBlackTree[T]) InOrderTraversal(callback func(val T)) {
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

func (rbt *redBlackTree[T]) PostOrderTraversal(callback func(val T)) {
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

func (rbt *redBlackTree[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.New[*rbtNode[T]]()
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
