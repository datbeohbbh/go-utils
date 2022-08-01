package linkedlist

func (li *linkedList[T]) empty() bool {
	return li.head == li.tail && li.head == nil
}
