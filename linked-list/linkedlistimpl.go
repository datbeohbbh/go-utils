package linkedlist

import (
	compare "github.com/datbeohbbh/go-utils/interfaces"
)

type linkedList[T compare.IEqual[T]] struct {
	head *node[T]
	tail *node[T]
	size int64
}

func New[T compare.IEqual[T]]() *linkedList[T] {
	return &linkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (li *linkedList[T]) Insert(value T) {
	li.insert(&node[T]{
		value: value,
		next:  nil,
	})
}

func (li *linkedList[T]) Empty() bool {
	return li.empty()
}

func (li *linkedList[T]) Remove(value T) bool {
	return li.remove(value)
}

func (li *linkedList[T]) Search(value T) bool {
	return li.search(value)
}

func (li *linkedList[T]) Size() int64 {
	return li.size
}
