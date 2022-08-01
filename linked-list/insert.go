package linkedlist

func (li *linkedList[T]) insert(item *node[T]) {
	if li.empty() {
		li.head = item
	} else {
		li.tail.next = item
	}
	li.tail = item
	li.size += 1
}
