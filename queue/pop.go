package queue

func (qu *queue[T]) Pop() {
	if !qu.Empty() {
		qu.head = qu.head.next
		qu.size -= 1
	}
}
