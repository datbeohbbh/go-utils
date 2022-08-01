package treap

func (curNode *node[K, V]) search(key K) (*node[K, V], bool) {
	if curNode.empty() {
		return nil, false
	}
	if curNode.key.Equal(key) {
		return curNode, true
	} else if curNode.key.Less(key) {
		return curNode.right.search(key)
	} else {
		return curNode.left.search(key)
	}
}
