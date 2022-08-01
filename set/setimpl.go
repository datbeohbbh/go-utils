package avltree

import (
	compare "github.com/datbeohbbh/go-utils/set/interfaces"
)

type treeSet[T compare.IComparator[T]] struct {
	tree *node[T]
}

func New[T compare.IComparator[T]]() *treeSet[T] {
	return &treeSet[T]{
		tree: nil,
	}
}

func (curTreeSet *treeSet[T]) Empty() bool {
	return curTreeSet.tree.empty()
}

func (curTreeSet *treeSet[T]) Insert(value T) bool {
	inserted := true
	curTreeSet.tree, inserted = curTreeSet.tree.insert(value)
	return inserted
}

func (curTreeSet *treeSet[T]) Remove(value T) bool {
	removed := true
	curTreeSet.tree, removed = curTreeSet.tree.remove(value)
	return removed
}

func (curTreeSet *treeSet[T]) Find(value T) bool {
	return (curTreeSet.tree.find(value))
}

func (curTreeSet *treeSet[T]) LowerBound(value T) (*T, bool) {
	lowerBoundNode := curTreeSet.tree.lowerBound(value)
	if lowerBoundNode == nil {
		return nil, false
	} else {
		return &lowerBoundNode.value, true
	}
}

func (curTreeSet *treeSet[T]) UpperBound(value T) (*T, bool) {
	upperBoundNode := curTreeSet.tree.upperBound(value)
	if upperBoundNode == nil {
		return nil, false
	} else {
		return &upperBoundNode.value, true
	}
}

func (curTreeSet *treeSet[T]) Size() int64 {
	if curTreeSet.Empty() {
		return 0
	}
	return curTreeSet.tree.size
}

func (curTreeSet *treeSet[T]) Clear() {
	if !curTreeSet.Empty() {
		curTreeSet.tree = nil
	}
}

func (curTreeSet *treeSet[T]) Begin() *T {
	res := curTreeSet.tree.begin()
	if res == nil {
		return nil
	} else {
		return &(res.value)
	}
}

func (curTreeSet *treeSet[T]) End() *T {
	res := curTreeSet.tree.end()
	if res == nil {
		return nil
	} else {
		return &(res.value)
	}
}