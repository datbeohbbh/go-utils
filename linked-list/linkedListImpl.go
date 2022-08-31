package linkedlist

import (
	compare "github.com/datbeohbbh/go-utils/interfaces"
)

type LinkedList[T compare.IEqual[T]] struct {
	head *node[T]
	tail *node[T]
	size int64
}

func New[T compare.IEqual[T]]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}
