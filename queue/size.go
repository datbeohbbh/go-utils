package queue

func (qu *queue[T]) Size() int64 {
	if qu.Empty() {
		return 0
	} else {
		return qu.size
	}
}
