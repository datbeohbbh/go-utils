package treap

func (curNode *node[K, V]) lowerBound(key K) *node[K, V] {
	if curNode.empty() {
		return nil
	}
	if curNode.key.Equal(key) {
		return curNode
	} else if curNode.key.Less(key) {
		return curNode.right.lowerBound(key)
	} else {
		lowerNode := curNode.left.lowerBound(key)
		if lowerNode == nil {
			return curNode
		} else {
			return lowerNode
		}
	}
}
