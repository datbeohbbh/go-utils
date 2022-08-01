package avltree

func (curNode *node[T]) begin() *node[T] {
	if curNode.empty() {
		return nil
	}
	if !curNode.left.empty() {
		return curNode.left.begin()
	} else {
		return curNode
	}
}
