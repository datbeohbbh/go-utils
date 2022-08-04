package queue

func (qu *queue[T]) Empty() bool {
	return qu.head == qu.tail && qu.head == nil
}
