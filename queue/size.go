package queue

func (qu *queue[T]) Size() int64 {
	return qu.size
}
