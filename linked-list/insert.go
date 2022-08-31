package linkedlist

func (li *LinkedList[T]) Insert(value T) {
	item := &node[T]{
		value: value,
		next:  nil,
	}
	if li.Empty() {
		li.head = item
	} else {
		li.tail.next = item
	}
	li.tail = item
	li.size += 1
}
