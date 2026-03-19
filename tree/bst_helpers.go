package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type BinarySearchTree[T cmp.Ordered] interface {
	Empty() bool
	Contains(T) bool
	Min() T
	Max() T
	Height() int
	Insert(T)
	Delete(T)
	PreOrderTraversal(func(T))
	InOrderTraversal(func(T))
	PostOrderTraversal(func(T))
	LevelOrderTraversal(func(T))
}

type baseBSTNodeInterface[T cmp.Ordered] interface {
	getVal() T
	getLeft() baseBSTNodeInterface[T]
	getRight() baseBSTNodeInterface[T]
}

func contains[T cmp.Ordered](root baseBSTNodeInterface[T], val T) bool {
	if root == nil {
		return false
	}

	if val == root.getVal() {
		return true
	}

	if val < root.getVal() {
		return contains(root.getLeft(), val)
	}

	return contains(root.getRight(), val)
}

func minChild[T cmp.Ordered](root baseBSTNodeInterface[T]) T {
	if root == nil {
		var zero T
		return zero
	}

	for root.getLeft() != nil {
		root = root.getLeft()
	}

	return root.getVal()
}

func maxChild[T cmp.Ordered](root baseBSTNodeInterface[T]) T {
	if root == nil {
		var zero T
		return zero
	}

	for root.getRight() != nil {
		root = root.getRight()
	}

	return root.getVal()
}

func height[T cmp.Ordered](root baseBSTNodeInterface[T]) int {
	if root == nil {
		return 0
	}

	return max(height(root.getLeft()), height(root.getRight())) + 1
}

func preOrderTraversal[T cmp.Ordered](root baseBSTNodeInterface[T], fn func(T)) {
	if root == nil {
		return
	}

	fn(root.getVal())
	preOrderTraversal(root.getLeft(), fn)
	preOrderTraversal(root.getRight(), fn)
}

func inOrderTraversal[T cmp.Ordered](root baseBSTNodeInterface[T], fn func(T)) {
	if root == nil {
		return
	}

	inOrderTraversal(root.getLeft(), fn)
	fn(root.getVal())
	inOrderTraversal(root.getRight(), fn)
}

func postOrderTraversal[T cmp.Ordered](root baseBSTNodeInterface[T], fn func(T)) {
	if root == nil {
		return
	}

	postOrderTraversal(root.getLeft(), fn)
	postOrderTraversal(root.getRight(), fn)
	fn(root.getVal())
}

func levelOrderTraversal[T cmp.Ordered](root baseBSTNodeInterface[T], fn func(T)) {
	if root == nil {
		return
	}

	q := queue.New[baseBSTNodeInterface[T]]()
	q.Enqueue(root)

	for !q.Empty() {
		node := q.Deque()
		if node == nil {
			continue
		}

		if node.getLeft() != nil {
			q.Enqueue(node.getLeft())
		}

		if node.getRight() != nil {
			q.Enqueue(node.getRight())
		}

		fn(node.getVal())
	}
}
