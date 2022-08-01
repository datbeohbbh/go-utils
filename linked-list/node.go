package linkedlist

import compare "github.com/datbeohbbh/go-utils/interfaces"

type node[T compare.IEqual[T]] struct {
	value T
	next  *node[T]
}
