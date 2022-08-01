package treap

func (curNode *node[K, V]) erase(key K) (*node[K, V], bool) {
	erased := true
	if curNode.empty() {
		return nil, false
	}
	if curNode.key.Equal(key) {
		curNode = merge(curNode.left, curNode.right)
		erased = true
	} else if curNode.key.Less(key) {
		curNode.right, erased = curNode.right.erase(key)
	} else {
		curNode.left, erased = curNode.left.erase(key)
	}
	curNode.update()
	return curNode, erased
}
