package linkedlist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/datbeohbbh/go-utils/types"
)


func TestInsert(t *testing.T) {
	li := New[Integer32]()

	a := []int{1, 2, 3, 4, 6, 6, 5, 4, 3, 2, 1}
	for _, v := range a {
		vv := types.NewInteger(int32(v))
		li.Insert(vv)
	}

	if li.Size() != 11 {
		t.Errorf("expected: %d, found: %d", 11, li.Size())
	}
}

func TestInsertStress(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	li := New[Integer32]()

	n := 1000000
	l := 1
	r := rand.Intn(1000000000)

	for i := 1; i <= n; i++ {
		v := types.NewInteger(int32(l + rand.Intn(r-l+1)))
		li.Insert(v)
	}

	if li.Size() != int64(n) {
		t.Errorf("expected: %d, found: %d", n, li.Size())
	}
}
