package avltree

import (
	compare "github.com/datbeohbbh/go-utils/set/interfaces"
)

type node[T compare.IComparator] struct {
	height  int64
	size    int64
	balance int64
	value   T
	left    *node[T]
	right   *node[T]
}
