package avltree

func (curNode *node[T]) last() *node[T] {
	if curNode.empty() {
		return nil
	}
	if !curNode.right.empty() {
		return curNode.right.last()
	} else {
		return curNode
	}
}
