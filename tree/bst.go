package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type BinarySearchTree[T cmp.Ordered] struct {
	root *binaryNode[T]
}

type binaryNode[T cmp.Ordered] struct {
	val   T
	left  *binaryNode[T]
	right *binaryNode[T]
}

func (bst *BinarySearchTree[T]) Empty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree[T]) Contains(val T) bool {
	return bst.root.contains(val)
}

func (n *binaryNode[T]) contains(val T) bool {
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

func (bst *BinarySearchTree[T]) Insert(val T) {
	bst.root = bst.root.insert(val)
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

func (bst *BinarySearchTree[T]) Delete(val T) {
	bst.root = bst.root.delete(val)
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

func (bst *BinarySearchTree[T]) PreOrderTraversal(callback func(val T)) {
	bst.root.preOrderTraversal(callback)
}

func (n *binaryNode[T]) preOrderTraversal(callback func(val T)) {
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

func (n *binaryNode[T]) inOrderTraversal(callback func(val T)) {
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

func (n *binaryNode[T]) postOrderTraversal(callback func(val T)) {
	if n == nil {
		return
	}

	n.left.postOrderTraversal(callback)
	n.right.postOrderTraversal(callback)
	callback(n.val)
}

func (bst *BinarySearchTree[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.Queue[*binaryNode[T]]{}
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
