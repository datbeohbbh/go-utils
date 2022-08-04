package queue

func (qu *queue[T]) Front() *T {
	if qu.Empty() {
		return nil
	}
	return &qu.head.value
}
