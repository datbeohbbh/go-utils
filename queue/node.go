package queue

type node[T any] struct {
	next  *node[T]
	value T
}
