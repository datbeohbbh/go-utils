package queue

func (qu *Queue[T]) Pop() {
	if !qu.Empty() {
		qu.head = qu.head.next
		qu.size -= 1
	}
}
