package numbers

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

func NewFloat[T interfaces.IFloat](value T) Float[T] {
	var v Float[T]
	v.SetValue(value)
	return v
}

type Float32 = Float[float32]

func NewFloat32(value float32) Float32 {
	return NewFloat(value)
}

type Float64 = Float[float64]

func NewFloat64(value float64) Float64 {
	return NewFloat(value)
}
