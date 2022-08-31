package linkedlist

func (li *linkedList[T]) GetTail() *T {
	if li.Empty() {
		return nil
	}
	return &li.tail.value
}
