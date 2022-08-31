package queue

func (qu *Queue[T]) Empty() bool {
	return qu.Size() == 0
}
