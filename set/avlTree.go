package avltree

import (
	compare "github.com/datbeohbbh/go-utils/interfaces"
)

type node[T compare.IComparator[T]] struct {
	height  int64
	size    int64
	balance int64
	value   T
	left    *node[T]
	right   *node[T]
}
