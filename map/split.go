package treap

func (curNode *node[K, V]) split(key K) (*node[K, V], *node[K, V]) {
	if curNode.empty() {
		return nil, nil
	}
	var leftNode *node[K, V]
	var rightNode *node[K, V]
	if curNode.key.Less(key) {
		leftNode = curNode
		leftNode.right, rightNode = curNode.right.split(key)
	} else {
		rightNode = curNode
		leftNode, rightNode.left = curNode.left.split(key)
	}
	leftNode.update()
	rightNode.update()
	return leftNode, rightNode
}
