package avltree

import "github.com/datbeohbbh/go-utils/types"

type T = types.Integer32

type Pair struct {
	first  int
	second int
}

func (t Pair) Less(rhs Pair) bool {
	return (t.first < rhs.first) || (t.first == rhs.first && t.second < rhs.second)
}

func (t Pair) Equal(rhs Pair) bool {
	return (t.first == rhs.first && t.second == rhs.second)
}
