package avltree

import (
	compare "github.com/datbeohbbh/go-utils/set/interfaces"
)

type set[T compare.IComparator] struct {
	tree *node[T]
}

func New[T compare.IComparator]() *set[T] {
	return &set[T]{
		tree: nil,
	}
}

func (curSet *set[T]) Empty() bool {
	return curSet.tree.empty()
}

func (curSet *set[T]) Insert(value T) bool {
	inserted := true
	curSet.tree, inserted = curSet.tree.insert(value)
	return inserted
}

func (curSet *set[T]) Remove(value T) bool {
	removed := true
	curSet.tree, removed = curSet.tree.remove(value)
	return removed
}

func (curSet *set[T]) Find(value T) bool {
	return (curSet.tree.find(value))
}

func (curSet *set[T]) LowerBound(value T) (*T, bool) {
	lowerBoundNode := curSet.tree.lowerBound(value)
	if lowerBoundNode == nil {
		return nil, false
	} else {
		return &lowerBoundNode.value, true
	}
}

func (curSet *set[T]) UpperBound(value T) (*T, bool) {
	upperBoundNode := curSet.tree.upperBound(value)
	if upperBoundNode == nil {
		return nil, false
	} else {
		return &upperBoundNode.value, true
	}
}

func (curSet *set[T]) Size() int64 {
	if curSet.Empty() {
		return 0
	}
	return curSet.tree.size
}

func (curSet *set[T]) Clear() {
	if !curSet.Empty() {
		curSet.tree = nil
	}
}
