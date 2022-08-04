package queue

type queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int64
}

func New[T any]() *queue[T] {
	return &queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}
