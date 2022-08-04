package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestPop(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	qu := New[int]()

	n := 1000000
	a := []int{}
	for i := 1; i <= n; i++ {
		val := 1 + rand.Intn(1000000000)
		a = append(a, val)
		qu.Push(val)
	}

	for i := 0; i < n-1; i++ {
		qu.Pop()
		if *qu.Front() != a[i+1] {
			t.Errorf("expected: %d, found: %d", a[i+1], *qu.Front())
		}
	}
}
