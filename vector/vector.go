package vector

import (
	"fmt"

	compare "github.com/datbeohbbh/go-utils/interfaces"
)

type Vector[T compare.IComparator[T]] struct {
	size int
	data []T
}

func New[T compare.IComparator[T]]() Vector[T] {
	return Vector[T]{
		size: 0,
		data: make([]T, 0), // len = 0, cap = 0
	}
}

func (v *Vector[T]) Size() int {
	// debug purpose
	if v.size != len(v.data) {
		panic(fmt.Errorf("expect: v.size should be equal to len(v.data) - actual v.size = %d, len(v.data) = %d", v.size, len(v.data)))
	}
	return v.size
}

func (v *Vector[T]) Get(index int) T {
	if index < 0 || index >= v.Size() {
		panic(fmt.Errorf("index %d out of range [0, %d)", index, v.Size()))
	}
	return v.data[index]
}

func (v *Vector[T]) Set(index int, value T) {
	if index < 0 || index >= v.Size() {
		panic(fmt.Errorf("index %d out of range [0, %d)", index, v.Size()))
	}
	v.data[index] = value
}

func (v *Vector[T]) Back() T {
	return v.Get(v.Size() - 1)
}

func insert[T compare.IComparator[T]](data []T, pos int, value T) []T {
	L := len(data)

	if pos < 0 || pos > L {
		panic(fmt.Errorf("index %d out of range [0, %d)", pos, L))
	}

	ret := data
	if cap(data) == L {
		ret = make([]T, L+1, (L+1)<<1)
		copy(ret, data[:pos])
	} else {
		ret = append(ret, value) // append dummy value
	}
	copy(ret[pos+1:], data[pos:])
	ret[pos] = value
	return ret
}

func (v *Vector[T]) Insert(pos int, value T) {
	v.data = insert(v.data, pos, value)
	v.size += 1
}

func (v *Vector[T]) PushBack(value T) {
	v.Insert(v.Size(), value)
}

func remove[T compare.IComparator[T]](data []T, pos int) []T {
	return append(data[:pos], data[pos+1:]...)
}

func (v *Vector[T]) Remove(index int) {
	if index < 0 || index >= v.Size() {
		panic(fmt.Errorf("index %d out of range [0, %d)", index, v.Size()))
	}
	v.data = remove(v.data, index)
	v.size -= 1
}

func (v *Vector[T]) PopBack() {
	v.Remove(v.Size() - 1)
}

func (v *Vector[T]) IndexOf(value T) int {
	for i := 0; i < v.Size(); i++ {
		if v.Get(i).Equal(value) {
			return i
		}
	}
	return -1
}

func (v *Vector[T]) BinarySearch(value T) int {
	lo, hi := 0, v.Size()-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if v.Get(mid).Equal(value) {
			return mid
		} else if v.Get(mid).Less(value) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// smallest greater or equal to value
func (v *Vector[T]) LowerBound(value T) int {
	lo, hi := 0, v.Size()-1
	pos := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if v.Get(mid).Less(value) {
			lo = mid + 1
		} else {
			hi = mid - 1
			pos = mid
		}
	}
	return pos
}

// smallest greater than value
func (v *Vector[T]) UpperBound(value T) int {
	lo, hi := 0, v.Size()-1
	pos := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if v.Get(mid).Less(value) || v.Get(mid).Equal(value) {
			lo = mid + 1
		} else {
			hi = mid - 1
			pos = mid
		}
	}
	return pos
}
