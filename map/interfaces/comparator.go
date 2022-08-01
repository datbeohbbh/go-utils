package interfaces

type IComparator[K any] interface {
	Less(K) bool
	Equal(K) bool
}
