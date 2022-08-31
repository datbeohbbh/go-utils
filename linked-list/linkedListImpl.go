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
