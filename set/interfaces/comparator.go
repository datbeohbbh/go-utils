package interfaces

type IComparator interface {
	Less(any) bool
	Equal(any) bool
}
