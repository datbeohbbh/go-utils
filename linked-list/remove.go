package linkedlist

// Remove by value
// Remove first position which equal to T.
func (li *linkedList[T]) Remove(value T) bool {
	found := li.Search(value)
	if !found {
		return false
	}
	if li.size == 1 {
		li.head = nil
		li.tail = nil
	} else {
		cur := li.head
		var prev *node[T] = nil
		var next *node[T] = nil
		for cur != nil {
			if cur.value.Equal(value) {
				next = cur.next
				break
			}
			prev = cur
			cur = cur.next
		}
		if prev == nil {
			li.head = next
		} else if next == nil {
			li.tail = prev
			li.tail.next = nil
		} else {
			prev.next = next
		}
	}
	li.size -= 1
	return true
}
