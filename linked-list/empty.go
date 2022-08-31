package linkedlist

func (li *LinkedList[T]) Empty() bool {
	return li.Size() == 0
}
