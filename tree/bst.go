package tree

import "cmp"

func NewBinarySearchTree[T cmp.Ordered]() BinarySearchTree[T] {
	return &binarySearchTree[T]{}
}

type binarySearchTree[T cmp.Ordered] struct {
	root *bstNode[T]
	baseBST[T]
}

type bstNode[T cmp.Ordered] struct {
	val   T
	left  *bstNode[T]
	right *bstNode[T]
}

func (n *bstNode[T]) getVal() T {
	return n.val
}

func (n *bstNode[T]) getLeft() baseBSTNodeInterface[T] {
	if n.left == nil {
		return nil
	}

	return n.left
}

func (n *bstNode[T]) getRight() baseBSTNodeInterface[T] {
	if n.right == nil {
		return nil
	}

	return n.right
}

func (bst *binarySearchTree[T]) Insert(val T) {
	bst.root = bst.root.insert(val)
	bst.baseBST.root = bst.root
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

func (bst *binarySearchTree[T]) Delete(val T) {
	bst.root = bst.root.delete(val)
	bst.baseBST.root = bst.root
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

	successorVal := minChild(n.right)
	n.val = successorVal
	n.right = n.right.delete(successorVal)
	return n
}
