package types

import (
	"github.com/datbeohbbh/go-utils/interfaces"
)

type Integer[T interfaces.IInteger] struct {
	interfaces.IComparator[T]
	value T
}

func (t Integer[T]) Less(rhs Integer[T]) bool {
	return (t.value < rhs.value)
}

func (t Integer[T]) Equal(rhs Integer[T]) bool {
	return (t.value == rhs.value)
}

func (t *Integer[T]) SetValue(value T) {
	t.value = value
}

func (t *Integer[T]) GetValue() T {
	return t.value
}

func NewInteger[T interfaces.IInteger](value T) Integer[T] {
	var v Integer[T]
	v.SetValue(value)
	return v
}

type Integer8 = Integer[int8]

type Integer16 = Integer[int16]

type Integer32 = Integer[int32]

type Integer64 = Integer[int64]
