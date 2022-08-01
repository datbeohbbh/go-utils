package treap

func (curNode *node[K, V]) empty() bool {
	return (curNode == nil || curNode.size == 0)
}
