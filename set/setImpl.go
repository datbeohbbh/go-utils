package avltree

import (
	compare "github.com/datbeohbbh/go-utils/interfaces"
)

type TreeSet[T compare.IComparator[T]] struct {
	tree *node[T]
}

func New[T compare.IComparator[T]]() *TreeSet[T] {
	return &TreeSet[T]{
		tree: nil,
	}
}

func (curTreeSet *TreeSet[T]) Empty() bool {
	return curTreeSet.tree.empty()
}

func (curTreeSet *TreeSet[T]) Insert(value T) bool {
	inserted := true
	curTreeSet.tree, inserted = curTreeSet.tree.insert(value)
	return inserted
}

func (curTreeSet *TreeSet[T]) Remove(value T) bool {
	removed := true
	curTreeSet.tree, removed = curTreeSet.tree.remove(value)
	return removed
}

func (curTreeSet *TreeSet[T]) Find(value T) bool {
	return (curTreeSet.tree.find(value))
}

func (curTreeSet *TreeSet[T]) LowerBound(value T) (*T, bool) {
	lowerBoundNode := curTreeSet.tree.lowerBound(value)
	if lowerBoundNode == nil {
		return nil, false
	} else {
		return &lowerBoundNode.value, true
	}
}

func (curTreeSet *TreeSet[T]) UpperBound(value T) (*T, bool) {
	upperBoundNode := curTreeSet.tree.upperBound(value)
	if upperBoundNode == nil {
		return nil, false
	} else {
		return &upperBoundNode.value, true
	}
}

func (curTreeSet *TreeSet[T]) Size() int64 {
	if curTreeSet.Empty() {
		return 0
	}
	return curTreeSet.tree.size
}

func (curTreeSet *TreeSet[T]) Clear() {
	if !curTreeSet.Empty() {
		curTreeSet.tree = nil
	}
}

func (curTreeSet *TreeSet[T]) Begin() *T {
	res := curTreeSet.tree.begin()
	if res == nil {
		return nil
	} else {
		return &(res.value)
	}
}

func (curTreeSet *TreeSet[T]) Last() *T {
	res := curTreeSet.tree.last()
	if res == nil {
		return nil
	} else {
		return &(res.value)
	}
}

func (curTreeSet *TreeSet[T]) Iterator() func() *T {
	curValue := curTreeSet.Begin()
	return func() *T {
		if curValue == nil {
			return nil
		} else {
			retValue := curValue
			curValue, _ = curTreeSet.UpperBound(*curValue)
			return retValue
		}
	}
}
