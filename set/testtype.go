package avltree

import "log"

type T struct {
	value int
}

func (t T) Cast(rhs any) T {
	x, same := rhs.(T)
	if !same {
		log.Panic("type error!!")
	}
	return x
}

func (t T) Less(rhs any) bool {
	return t.value < t.Cast(rhs).value
}

func (t T) Equal(rhs any) bool {
	return (t.value == t.Cast(rhs).value)
}

type Pair struct {
	first  int
	second int
}

func (t Pair) Cast(rhs any) Pair {
	x, same := rhs.(Pair)
	if !same {
		log.Panic("type error!!")
	}
	return x
}

func (t Pair) Less(rhs any) bool {
	x := t.Cast(rhs)
	return (t.first < x.first) || (t.first == x.first && t.second < x.second)
}

func (t Pair) Equal(rhs any) bool {
	x := t.Cast(rhs)
	return (t.first == x.first && t.second == x.second)
}
