package vector

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

type Pair struct {
	key   []byte
	value int
}

func (p *Pair) Less(other *Pair) bool {
	return bytes.Compare(p.key, other.key) == -1
}

func (p *Pair) Equal(other *Pair) bool {
	return bytes.Equal(p.key, other.key)
}

func TestVectorPushBack(t *testing.T) {
	var v Vector[*Pair]

	for i := 1; i <= 10; i++ {
		v.PushBack(&Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
	}

	if v.Size() != 10 {
		t.Errorf("expected size = %d, actual size = %d", 10, v.Size())
	}

	for i := 0; i < v.Size(); i++ {
		log.Println(string(v.Get(i).key), v.Get(i).value, v.Size())
	}
}

func TestVectorInsert(t *testing.T) {
	v := New[*Pair]()
	for i := 1; i <= 10; i++ {
		v.Insert(0, &Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
	}

	// should print in reverse order
	for i := 0; i < v.Size(); i++ {
		log.Println(string(v.Get(i).key), v.Get(i).value, v.Size())
	}

	v.Insert(5, &Pair{
		key:   []byte(fmt.Sprintf("insert at %d", 5)),
		value: -1,
	})

	for i := 0; i < v.Size(); i++ {
		log.Println(string(v.Get(i).key), v.Get(i).value, v.Size())
	}
}

func TestRemove(t *testing.T) {
	v := New[*Pair]()
	for i := 1; i <= 10; i++ {
		v.Insert(0, &Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
	}

	v.Remove(5)

	for i := 0; i < v.Size(); i++ {
		log.Println(string(v.Get(i).key), v.Get(i).value, v.Size())
	}
}

func TestPopBack(t *testing.T) {
	v := New[*Pair]()
	for i := 1; i <= 10; i++ {
		v.Insert(0, &Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
	}

	for i := 1; i <= 5; i++ {
		v.PopBack()
	}

	for i := 0; i < v.Size(); i++ {
		log.Println(string(v.Get(i).key), v.Get(i).value, v.Size())
	}
}

func TestLinearSearch(t *testing.T) {
	v := New[*Pair]()
	for i := 1; i <= 10; i++ {
		v.PushBack(&Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
	}

	for i := 1; i <= 10; i++ {
		pos := v.IndexOf(&Pair{
			key:   []byte(fmt.Sprintf("Hello %d", i)),
			value: i,
		})
		if pos == -1 {
			t.Errorf("expected found pos = %d, actual pos = %d", i-1, pos)
		}
	}
}

type num struct {
	value int
}

func (n num) Less(other num) bool {
	return n.value < other.value
}

func (n num) Equal(other num) bool {
	return n.value == other.value
}

func TestBinarySearch(t *testing.T) {
	v := New[num]()

	for i := 1; i <= 100000; i++ {
		v.PushBack(num{value: i})
	}

	for i := 1; i <= 100000; i++ {
		pos := v.BinarySearch(num{value: i})
		if pos != i-1 {
			t.Errorf("expected found pos = %d, actual pos = %d", i-1, pos)
		}
	}
}

func TestLowerBound(t *testing.T) {
	v := New[num]()

	for i := 1; i <= 100000; i++ {
		v.PushBack(num{value: i})
	}

	for i := 1; i <= 100000; i++ {
		pos := v.LowerBound(num{value: i})
		if pos != i-1 || v.Get(pos).value != i {
			t.Errorf("expected found pos = %d, actual pos = %d", i-1, pos)
		}
	}
}

func TestUpperBound(t *testing.T) {
	v := New[num]()

	for i := 1; i <= 100000; i++ {
		v.PushBack(num{value: i})
	}

	for i := 1; i <= 99999; i++ {
		pos := v.UpperBound(num{value: i})
		if pos != i || v.Get(pos).value != i+1 {
			t.Errorf("expected found pos = %d, actual pos = %d", i, pos)
		}
	}

	pos := v.UpperBound(num{value: 100000})
	if pos != -1 {
		t.Errorf("expected found pos = %d, actual pos = %d", -1, pos)
	}
}
