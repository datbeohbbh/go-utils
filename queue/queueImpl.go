package queue

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int64
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}
