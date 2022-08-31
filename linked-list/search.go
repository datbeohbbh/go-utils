package linkedlist

func (li *LinkedList[T]) Search(value T) bool {
	if li.Empty() {
		return false
	}
	head := li.head
	for head != nil {
		if head.value.Equal(value) {
			return true
		}
		head = head.next
	}
	return false
}
