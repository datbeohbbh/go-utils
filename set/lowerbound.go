package avltree

func (curNode *node[T]) lowerBound(value T) *node[T] {
	if curNode.empty() {
		return nil
	}
	if curNode.value.Equal(value) {
		return curNode
	} else if curNode.value.Less(value) {
		return curNode.right.lowerBound(value)
	} else {
		lowerNode := curNode.left.lowerBound(value)
		if lowerNode == nil {
			return curNode
		} else {
			return lowerNode
		}
	}
}
