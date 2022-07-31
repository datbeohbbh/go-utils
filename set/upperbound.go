package avltree

func (curNode *node[T]) upperBound(value T) *node[T] {
	if curNode.empty() {
		return nil
	}
	if curNode.value.Equal(value) || curNode.value.Less(value) {
		return curNode.right.upperBound(value)
	} else {
		upperNode := curNode.left.upperBound(value)
		if upperNode == nil {
			return curNode
		} else {
			return upperNode
		}
	}
}
