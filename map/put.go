package treap

func (curNode *node[K, V]) put(item *node[K, V]) *node[K, V] {
	if curNode.empty() {
		return item
	} else if curNode.priority < item.priority {
		item.left, item.right = curNode.split(item.key)
		curNode = item
	} else {
		if curNode.key.Less(item.key) {
			curNode.right = curNode.right.put(item)
		} else {
			curNode.left = curNode.left.put(item)
		}
	}
	curNode.update()
	return curNode
}
