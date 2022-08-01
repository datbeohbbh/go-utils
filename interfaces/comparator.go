package interfaces

type IEqual[T any] interface {
	Equal(T) bool
}

type ILess[T any] interface {
	Less(T) bool
}

type IGreater[T any] interface {
	Greater(T) bool
}

type IComparator[T any] interface {
	IEqual[T]
	ILess[T]
}
