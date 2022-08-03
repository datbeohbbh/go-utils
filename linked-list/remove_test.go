package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

func removeIndex(arr []int, i int) []int {
	b := append(arr[:i], arr[i+1:]...)
	return b
}

func TestRemoveSmall(t *testing.T) {

	a := []int{4, 6, 21, 612, 56123, 65, 3, 256, 31, 3, 512, 312, 5, 5, 5, 6, 7, 1225661, 11}
	li := New[Integer32]()
	for _, v := range a {
		li.Insert(NewInteger(v))
	}
	removeIndexList := []int{0, 2, 2, 5, 8, 8, 12}
	for testID, idx := range removeIndexList {
		li.Remove(NewInteger(a[idx]))
		a = removeIndex(a, idx)
		b := []int{}
		for next := li.Iterator(); ; {
			val := next()
			if val == nil {
				break
			}
			b = append(b, int(val.GetValue()))
		}
		t.Logf("test #%d:\n%v, len: %d\n,%v, len: %d", testID, a, len(a), b, len(b))
		for next, c := li.Iterator(), 0; c < len(a); c++ {
			val := next()
			if val == nil || val.GetValue() != int32(a[c]) {
				if val == nil {
					t.Errorf("expected: %d, found: nil", a[c])
				} else {
					t.Errorf("expected: %d, found: %d", a[c], val.GetValue())
				}
			}
		}
	}
}

func TestRemoveStress(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	li := New[Integer32]()
	n := 10000
	l := 100
	r := 100000
	a := []int{}
	has := make(map[int]bool)

	for i := 1; i <= n; i++ {
		val := l + rand.Intn(r-l+1)
		for has[val] {
			val = l + rand.Intn(r-l+1)
		}
		a = append(a, val)
		has[val] = true
		li.Insert(NewInteger(val))
	}

	for len(a) > 0 {
		idx := rand.Intn(len(a))
		li.Remove(NewInteger(a[idx]))
		a = removeIndex(a, idx)
		for next, c := li.Iterator(), 0; c < len(a); c++ {
			val := next()
			if val == nil || val.GetValue() != int32(a[c]) {
				if val == nil {
					t.Errorf("expected: %d, found: nil", a[c])
				} else {
					t.Errorf("expected: %d, found: %d", a[c], val.GetValue())
				}
			}
		}
	}
}
