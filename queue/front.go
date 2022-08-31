package queue

func (qu *Queue[T]) Front() *T {
	if qu.Empty() {
		return nil
	}
	return &qu.head.value
}
