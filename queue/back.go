package queue

func (qu *Queue[T]) Back() *T {
	if qu.Empty() {
		return nil
	} else {
		return &qu.tail.value
	}
}
