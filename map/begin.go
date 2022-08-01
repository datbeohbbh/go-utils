package treap

func (curNode *node[K, V]) begin() *node[K, V] {
	if curNode.empty() {
		return nil
	}
	if curNode.left.empty() {
		return curNode
	} else {
		return curNode.left.begin()
	}
}
