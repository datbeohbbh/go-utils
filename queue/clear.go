package queue

func (qu *Queue[T]) Clear() {
	if !qu.Empty() {
		*qu = Queue[T]{}
	}
}
