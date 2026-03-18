package tree

import "cmp"

type BinaryTree[T cmp.Ordered] struct {
	root *binaryNode[T]
}

type binaryNode[T cmp.Ordered] struct {
	val   T
	left  *binaryNode[T]
	right *binaryNode[T]
}

func (bt *BinaryTree[T]) Empty() bool {
	return bt.root == nil
}

func (bt *BinaryTree[T]) Insert(val T) {
	bt.root = bt.root.insert(val)
}

func (n *binaryNode[T]) insert(val T) *binaryNode[T] {
	if n == nil {
		return &binaryNode[T]{val: val}
	}

	if val < n.val {
		n.left = n.left.insert(val)
	} else if val > n.val {
		n.right = n.right.insert(val)
	}
	return n
}

func (bt *BinaryTree[T]) Delete(val T) {
	bt.root = bt.root.delete(val)
}

func (n *binaryNode[T]) delete(val T) *binaryNode[T] {
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

	currentNode := n.right
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	n.val = currentNode.val
	n.right = n.right.delete(currentNode.val)
	return n
}

func (bt *BinaryTree[T]) PreOrderTraversal(callback func(val T)) {
	bt.root.preOrderTraversal(callback)
}

func (n *binaryNode[T]) preOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	callback(n.val)
	n.left.preOrderTraversal(callback)
	n.right.preOrderTraversal(callback)
}

func (bt *BinaryTree[T]) InOrderTraversal(callback func(val T)) {
	bt.root.inOrderTraversal(callback)
}

func (n *binaryNode[T]) inOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.inOrderTraversal(callback)
	callback(n.val)
	n.right.inOrderTraversal(callback)
}

func (bt *BinaryTree[T]) PostOrderTraversal(callback func(val T)) {
	bt.root.postOrderTraversal(callback)
}

func (n *binaryNode[T]) postOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.postOrderTraversal(callback)
	n.right.postOrderTraversal(callback)
	callback(n.val)
}
