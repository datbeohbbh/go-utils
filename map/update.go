package treap

func (curNode *node[K, V]) update() {
	if curNode.empty() {
		return
	}
	var lhs, rhs int64 = 0, 0
	if !curNode.left.empty() {
		lhs = curNode.left.size
	}
	if !curNode.right.empty() {
		rhs = curNode.right.size
	}
	curNode.size = lhs + 1 + rhs
}
