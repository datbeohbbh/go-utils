package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestSize(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	qu := New[int]()

	for rep := 1; rep <= 10; rep++ {
		n := 40000 + rand.Intn(100000)
		qu.Clear()
		for i := 1; i <= n; i++ {
			qu.Push(i)
		}
		if qu.Size() != int64(n) {
			t.Errorf("rep #%d - expected: %d, found: %d", rep, n, qu.Size())
		}
	}
}
