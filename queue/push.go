package queue

func (qu *Queue[T]) Push(value T) {
	item := &node[T]{
		next:  nil,
		value: value,
	}
	if qu.Empty() {
		qu.head = item
	} else {
		qu.tail.next = item
	}
	qu.tail = item
	qu.size += 1
}
