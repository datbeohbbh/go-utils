package queue

func (qu *queue[T]) Clear() {
	if !qu.Empty() {
		*qu = queue[T]{}
	}
}
