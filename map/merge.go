package treap

import compare "github.com/datbeohbbh/go-utils/map/interfaces"

func merge[K compare.IComparator[K], V any](leftNode *node[K, V], rightNode *node[K, V]) *node[K, V] {
	if leftNode.empty() || rightNode.empty() {
		if !leftNode.empty() {
			return leftNode
		} else {
			return rightNode
		}
	}
	if leftNode.priority > rightNode.priority {
		leftNode.right = merge(leftNode.right, rightNode)
		leftNode.update()
		return leftNode
	} else {
		rightNode.left = merge(leftNode, rightNode.left)
		rightNode.update()
		return rightNode
	}
}
