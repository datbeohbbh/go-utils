package avltree

func (curNode *node[T]) remove(value T) (*node[T], bool) {
	removed := true
	if curNode.empty() {
		return nil, false
	}
	if value.Equal(curNode.value) {
		if !(curNode.left.empty()) && !(curNode.right.empty()) {
			t := curNode.right.findLeftMost()
			curNode.value = t.value
			curNode.right, removed = curNode.right.remove(t.value)
		} else if !(curNode.left.empty()) {
			curNode = curNode.left
		} else if !(curNode.right.empty()) {
			curNode = curNode.right
		} else {
			curNode = nil
		}
	} else if value.Less(curNode.value) {
		curNode.left, removed = curNode.left.remove(value)
	} else {
		curNode.right, removed = curNode.right.remove(value)
	}
	if curNode.empty() {
		return nil, removed
	}
	curNode = curNode.balanceTree()
	return curNode, removed
}
