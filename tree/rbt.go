package tree

import "cmp"

func NewRedBlackTree[T cmp.Ordered]() BinarySearchTree[T] {
	return &redBlackTree[T]{}
}

type redBlackTree[T cmp.Ordered] struct {
	root *rbtNode[T]
}

func (t *redBlackTree[T]) Empty() bool                    { return t.root == nil }
func (t *redBlackTree[T]) Contains(val T) bool            { return contains(t.root, val) }
func (t *redBlackTree[T]) Min() T                         { return minChild(t.root) }
func (t *redBlackTree[T]) Max() T                         { return maxChild(t.root) }
func (t *redBlackTree[T]) Height() int                    { return height(t.root) }
func (t *redBlackTree[T]) PreOrderTraversal(fn func(T))   { preOrderTraversal(t.root, fn) }
func (t *redBlackTree[T]) InOrderTraversal(fn func(T))    { inOrderTraversal(t.root, fn) }
func (t *redBlackTree[T]) PostOrderTraversal(fn func(T))  { postOrderTraversal(t.root, fn) }
func (t *redBlackTree[T]) LevelOrderTraversal(fn func(T)) { levelOrderTraversal(t.root, fn) }

type rbtNode[T cmp.Ordered] struct {
	val    T
	left   *rbtNode[T]
	right  *rbtNode[T]
	parent *rbtNode[T]
	red    bool
}

func (n *rbtNode[T]) getVal() T {
	return n.val
}

func (n *rbtNode[T]) getLeft() baseBSTNodeInterface[T] {
	if n.left == nil {
		return nil
	}

	return n.left
}

func (n *rbtNode[T]) getRight() baseBSTNodeInterface[T] {
	if n.right == nil {
		return nil
	}

	return n.right
}

func (n *rbtNode[T]) isRed() bool {
	if n == nil {
		return false
	}

	return n.red
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
