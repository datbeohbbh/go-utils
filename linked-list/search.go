package linkedlist

func (li *linkedList[T]) search(value T) bool {
	if li.empty() {
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
