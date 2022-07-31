package avltree

func (curNode *node[T]) insert(value T) (*node[T], bool) {
	inserted := true
	if curNode.empty() {
		curNode = &node[T]{
			height:  0,
			size:    1,
			balance: 0,
			value:   value,
			left:    nil,
			right:   nil,
		}
		return curNode, true
	}

	if value.Equal(curNode.value) {
		return curNode, false
	} else if value.Less(curNode.value) {
		curNode.left, inserted = curNode.left.insert(value)
	} else {
		curNode.right, inserted = curNode.right.insert(value)
	}
	curNode = curNode.balanceTree()
	return curNode, inserted
}
