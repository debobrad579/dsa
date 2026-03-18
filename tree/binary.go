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
	if bt.root == nil {
		bt.root = &binaryNode[T]{val: val}
	} else {
		bt.root.insert(val)
	}
}

func (n *binaryNode[T]) insert(val T) {
	if val == n.val {
		return
	}

	if val < n.val {
		if n.left == nil {
			n.left = &binaryNode[T]{val: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &binaryNode[T]{val: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (bt *BinaryTree[T]) PreOrderTraversal(callback func(val T)) {
	if bt.root != nil {
		bt.root.preOrderTraversal(callback)
	}
}

func (n *binaryNode[T]) preOrderTraversal(callback func(val T)) {
	callback(n.val)

	if n.left != nil {
		n.left.preOrderTraversal(callback)
	}

	if n.right != nil {
		n.right.preOrderTraversal(callback)
	}
}

func (bt *BinaryTree[T]) InOrderTraversal(callback func(val T)) {
	if bt.root != nil {
		bt.root.inOrderTraversal(callback)
	}
}

func (n *binaryNode[T]) inOrderTraversal(callback func(val T)) {
	if n.left != nil {
		n.left.inOrderTraversal(callback)
	}

	callback(n.val)

	if n.right != nil {
		n.right.inOrderTraversal(callback)
	}
}

func (bt *BinaryTree[T]) PostOrderTraversal(callback func(val T)) {
	if bt.root != nil {
		bt.root.postOrderTraversal(callback)
	}
}

func (n *binaryNode[T]) postOrderTraversal(callback func(val T)) {
	if n.left != nil {
		n.left.postOrderTraversal(callback)
	}

	if n.right != nil {
		n.right.postOrderTraversal(callback)
	}

	callback(n.val)
}
