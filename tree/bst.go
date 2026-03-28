package tree

import (
	"cmp"
	"iter"
)

func NewBinarySearchTree[T cmp.Ordered]() BinarySearchTree[T] {
	return &binarySearchTree[T]{}
}

type binarySearchTree[T cmp.Ordered] struct {
	root *bstNode[T]
}

func (t *binarySearchTree[T]) Empty() bool         { return t.root == nil }
func (t *binarySearchTree[T]) Contains(val T) bool { return contains(t.root, val) }
func (t *binarySearchTree[T]) Min() T              { return minChild(t.root) }
func (t *binarySearchTree[T]) Max() T              { return maxChild(t.root) }
func (t *binarySearchTree[T]) Height() int         { return height(t.root) }

func (t *binarySearchTree[T]) PreOrderTraversal() iter.Seq[T] {
	return preOrderTraversal(t.root)
}

func (t *binarySearchTree[T]) InOrderTraversal() iter.Seq[T] {
	return inOrderTraversal(t.root)
}

func (t *binarySearchTree[T]) PostOrderTraversal() iter.Seq[T] {
	return postOrderTraversal(t.root)
}

func (t *binarySearchTree[T]) LevelOrderTraversal() iter.Seq[T] {
	return levelOrderTraversal(t.root)
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
