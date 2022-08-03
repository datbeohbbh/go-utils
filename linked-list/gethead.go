package linkedlist

func (li *linkedList[T]) GetHead() *T {
	if li.Empty() {
		return nil
	}
	return &li.head.value
}
