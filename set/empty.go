package avltree

func (curNode *node[T]) empty() bool {
	return (curNode == nil || curNode.size == 0)
}
