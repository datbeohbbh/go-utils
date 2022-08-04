package queue

func (qu *queue[T]) Back() *T {
	if qu.Empty() {
		return nil
	} else {
		return &qu.tail.value
	}
}
