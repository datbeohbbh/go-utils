package types

import (
	"github.com/datbeohbbh/go-utils/interfaces"
)

type Float[T interfaces.IFloat] struct {
	interfaces.IComparator[T]
	value T
}

func (t Float[T]) Less(rhs Float[T]) bool {
	return (t.value < rhs.value)
}

func (t Float[T]) Equal(rhs Float[T]) bool {
	return (t.value == rhs.value)
}

func (t *Float[T]) SetValue(value T) {
	t.value = value
}

func (t *Float[T]) GetValue() T {
	return t.value
}

type Float32 = Float[float32]

type Float64 = Float[float64]
