package queue

func (qu *Queue[T]) Size() int64 {
	return qu.size
}
