package avltree

func (curNode *node[T]) end() *node[T] {
	if curNode.empty() {
		return nil
	}
	if !curNode.right.empty() {
		return curNode.right.end()
	} else {
		return curNode
	}
}
