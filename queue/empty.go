package queue

func (qu *queue[T]) Empty() bool {
	return qu.Size() == 0
}
