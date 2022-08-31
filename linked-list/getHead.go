package linkedlist

func (li *LinkedList[T]) GetHead() *T {
	if li.Empty() {
		return nil
	}
	return &li.head.value
}
