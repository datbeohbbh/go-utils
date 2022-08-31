package treap

func (curNode *node[K, V]) upperBound(key K) *node[K, V] {
	if curNode.empty() {
		return nil
	}
	if curNode.key.Equal(key) || curNode.key.Less(key) {
		return curNode.right.upperBound(key)
	} else {
		upperNode := curNode.left.upperBound(key)
		if upperNode == nil {
			return curNode
		} else {
			return upperNode
		}
	}
}
