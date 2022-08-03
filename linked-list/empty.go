package linkedlist

func (li *linkedList[T]) Empty() bool {
	return li.head == li.tail && li.head == nil
}
