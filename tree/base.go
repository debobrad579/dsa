package tree

import (
	"cmp"

	"github.com/debobrad579/dsa/queue"
)

type BinarySearchTree[T cmp.Ordered] interface {
	Empty() bool
	Contains(T) bool
	Equals(baseBSTNodeInterface[T]) bool
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

type baseBST[T cmp.Ordered] struct {
	root baseBSTNodeInterface[T]
}

type baseBSTNodeInterface[T cmp.Ordered] interface {
	getVal() T
	getLeft() baseBSTNodeInterface[T]
	getRight() baseBSTNodeInterface[T]
}

func (bt *baseBST[T]) Empty() bool {
	return bt.root == nil
}

func (bt *baseBST[T]) Equals(other baseBSTNodeInterface[T]) bool {
	return equals(bt.root, other)
}

func equals[T cmp.Ordered](n baseBSTNodeInterface[T], other baseBSTNodeInterface[T]) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil || n.getVal() != other.getVal() {
		return false
	}

	return equals(n.getLeft(), other.getLeft()) && equals(n.getRight(), other.getRight())
}

func (bt *baseBST[T]) Contains(val T) bool {
	return contains(bt.root, val)
}

func contains[T cmp.Ordered](n baseBSTNodeInterface[T], val T) bool {
	if n == nil {
		return false
	}

	if val == n.getVal() {
		return true
	}

	if val < n.getVal() {
		return contains(n.getLeft(), val)
	}

	return contains(n.getRight(), val)
}

func (bt *baseBST[T]) Min() T {
	if bt.root == nil {
		var zero T
		return zero
	}

	return minChild(bt.root)
}

func minChild[T cmp.Ordered](n baseBSTNodeInterface[T]) T {
	if n.getLeft() == nil {
		return n.getVal()
	}

	return minChild(n.getLeft())
}

func (bt *baseBST[T]) Max() T {
	if bt.root == nil {
		var zero T
		return zero
	}

	return maxChild(bt.root)
}

func maxChild[T cmp.Ordered](n baseBSTNodeInterface[T]) T {
	if n.getRight() == nil {
		return n.getVal()
	}

	return maxChild(n.getRight())
}

func (bt *baseBST[T]) Height() int {
	return height(bt.root)
}

func height[T cmp.Ordered](n baseBSTNodeInterface[T]) int {
	if n == nil {
		return 0
	}

	return max(height(n.getLeft()), height(n.getRight())) + 1
}

func (bt *baseBST[T]) PreOrderTraversal(callback func(val T)) {
	preOrderTraversal(bt.root, callback)
}

func preOrderTraversal[T cmp.Ordered](n baseBSTNodeInterface[T], callback func(val T)) {
	if n == nil {
		return
	}

	callback(n.getVal())
	preOrderTraversal(n.getLeft(), callback)
	preOrderTraversal(n.getRight(), callback)
}

func (bt *baseBST[T]) InOrderTraversal(callback func(val T)) {
	inOrderTraversal(bt.root, callback)
}

func inOrderTraversal[T cmp.Ordered](n baseBSTNodeInterface[T], callback func(val T)) {
	if n == nil {
		return
	}

	inOrderTraversal(n.getLeft(), callback)
	callback(n.getVal())
	inOrderTraversal(n.getRight(), callback)
}

func (bt *baseBST[T]) PostOrderTraversal(callback func(val T)) {
	postOrderTraversal(bt.root, callback)
}

func postOrderTraversal[T cmp.Ordered](n baseBSTNodeInterface[T], callback func(val T)) {
	if n == nil {
		return
	}

	postOrderTraversal(n.getLeft(), callback)
	postOrderTraversal(n.getRight(), callback)
	callback(n.getVal())
}

func (bt *baseBST[T]) LevelOrderTraversal(callback func(val T)) {
	q := queue.New[baseBSTNodeInterface[T]]()
	q.Enqueue(bt.root)

	for !q.Empty() {
		n := q.Deque()
		if n == nil {
			continue
		}

		if n.getLeft() != nil {
			q.Enqueue(n.getLeft())
		}

		if n.getRight() != nil {
			q.Enqueue(n.getRight())
		}

		callback(n.getVal())
	}
}
