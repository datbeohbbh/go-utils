package linkedlist

func (li *LinkedList[T]) GetTail() *T {
	if li.Empty() {
		return nil
	}
	return &li.tail.value
}
