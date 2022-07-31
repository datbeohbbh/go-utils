package avltree

func (curNode *node[T]) find(value T) bool {
	if curNode.empty() {
		return false
	}
	findLeft, findRight := false, false
	if value.Equal(curNode.value) {
		return true
	} else if value.Less(curNode.value) {
		findLeft = curNode.left.find(value)
	} else {
		findRight = curNode.right.find(value)
	}
	return (findLeft || findRight)
}

func (curNode *node[T]) findLeftMost() *node[T] {
	t := curNode
	for t.left != nil {
		t = t.left
	}
	return t
}
