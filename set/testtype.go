package avltree

type T struct {
	value int
}

func (t T) Less(rhs T) bool {
	return t.value < rhs.value
}

func (t T) Equal(rhs T) bool {
	return (t.value == rhs.value)
}

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
