package linkedlist

func (li *linkedList[T]) Empty() bool {
	return li.Size() == 0
}
