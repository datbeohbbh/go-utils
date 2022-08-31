package numbers

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

func NewInteger8(value int8) Integer8 {
	return NewInteger(value)
}

type Integer16 = Integer[int16]

func NewInteger16(value int16) Integer16 {
	return NewInteger(value)
}

type Integer32 = Integer[int32]

func NewInteger32(value int) Integer32 {
	return NewInteger(int32(value))
}

type Integer64 = Integer[int64]

func NewInteger64(value int) Integer64 {
	return NewInteger(int64(value))
}
