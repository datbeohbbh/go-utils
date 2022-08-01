package interfaces

type IComparator[T any] interface {
	Less(T) bool
	Equal(T) bool
}
