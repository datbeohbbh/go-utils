package avltree

func (curNode *node[T]) reloadHeight() {
	var lhs, rhs int64
	if !curNode.left.empty() {
		lhs = curNode.left.height
	}
	if !curNode.right.empty() {
		rhs = curNode.right.height
	}
	if lhs < rhs {
		curNode.height = 1 + rhs
	} else {
		curNode.height = 1 + lhs
	}
}

func (curNode *node[T]) reloadBalance() {
	var lhs, rhs int64 = 0, 0
	if !curNode.left.empty() {
		lhs = curNode.left.height
	}
	if !curNode.right.empty() {
		rhs = curNode.right.height
	}
	curNode.balance = lhs - rhs
}

func (curNode *node[T]) reloadSize() {
	var lhs, rhs int64 = 0, 0
	if !curNode.left.empty() {
		lhs = curNode.left.size
	}
	if !curNode.right.empty() {
		rhs = curNode.right.size
	}
	curNode.size = 1 + lhs + rhs
}

func (curNode *node[T]) reload() {
	curNode.reloadHeight()
	curNode.reloadBalance()
	curNode.reloadSize()
}

func (curNode *node[T]) rotateLeft() *node[T] {
	newRoot := curNode.right
	leftRoot := newRoot.left
	newRoot.left = curNode
	curNode.right = leftRoot
	curNode.reload()
	newRoot.reload()
	return newRoot
}

func (curNode *node[T]) rotateRight() *node[T] {
	newRoot := curNode.left
	rightRoot := newRoot.right
	newRoot.right = curNode
	curNode.left = rightRoot
	curNode.reload()
	newRoot.reload()
	return newRoot
}

func (curNode *node[T]) rotateNode() *node[T] {
	if curNode.balance > 1 {
		if curNode.left.balance >= 0 {
			return curNode.rotateRight()
		} else {
			curNode.left = curNode.left.rotateLeft()
			return curNode.rotateRight()
		}
	}
	if curNode.balance < -1 {
		if curNode.right.balance <= 0 {
			return curNode.rotateLeft()
		} else {
			curNode.right = curNode.right.rotateRight()
			return curNode.rotateLeft()
		}
	}
	return curNode
}

func (curNode *node[T]) balanceTree() *node[T] {
	curNode.reload()
	return curNode.rotateNode()
}
