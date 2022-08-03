package treap

func (curNode *node[K, V]) last() *node[K, V] {
	if curNode.empty() {
		return nil
	}
	if curNode.right.empty() {
		return curNode
	} else {
		return curNode.right.last()
	}
}
