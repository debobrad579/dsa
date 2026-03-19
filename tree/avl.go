package tree

import "cmp"

func NewAVLTree[T cmp.Ordered]() BinarySearchTree[T] {
	return &avlTree[T]{}
}

type avlTree[T cmp.Ordered] struct {
	root *avlNode[T]
	baseBST[T]
}

type avlNode[T cmp.Ordered] struct {
	val    T
	left   *avlNode[T]
	right  *avlNode[T]
	height int
}

func (n *avlNode[T]) getVal() T {
	return n.val
}

func (n *avlNode[T]) getLeft() baseBSTNodeInterface[T] {
	if n.left == nil {
		return nil
	}

	return n.left
}

func (n *avlNode[T]) getRight() baseBSTNodeInterface[T] {
	if n.right == nil {
		return nil
	}

	return n.right
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

func (avl *avlTree[T]) Height() int {
	return avl.root.nodeHeight()
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

func (avl *avlTree[T]) Insert(val T) {
	avl.root = avl.root.insert(val)
	avl.baseBST.root = avl.root
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

func (avl *avlTree[T]) Delete(val T) {
	avl.root = avl.root.delete(val)
	avl.baseBST.root = avl.root
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
			successorVal := maxChild(n.left)
			n.val = successorVal
			n.left = n.left.delete(successorVal)
		} else {
			successorVal := minChild(n.right)
			n.val = successorVal
			n.right = n.right.delete(successorVal)
		}
	}

	n.updateHeight()
	return n.rotate()
}
