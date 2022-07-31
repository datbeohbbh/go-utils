package avltree

import "testing"

func TestUpperBoundSmall(t *testing.T) {
	set := New[T]()
	a := []int{4, 5, 6, 8, 9, 10, 13, 14, 15, 16, 17, 18, 20, 21}
	for _, v := range a {
		set.Insert(T{value: v})
	}

	res, _ := set.UpperBound(T{value: 9})
	if res == nil || (*res).value != 10 {
		t.Errorf("expected: %d, found %v", 10, res)
	}

	res, _ = set.UpperBound(T{value: 3})
	if res == nil || (*res).value != 4 {
		t.Errorf("expected: %d, found %v", 4, res)
	}

	res, _ = set.UpperBound(T{value: 7})
	if res == nil || (*res).value != 8 {
		t.Errorf("expected: %d, found %v", 8, res)
	}

	res, _ = set.UpperBound(T{value: 12})
	if res == nil || (*res).value != 13 {
		t.Errorf("expected: %d, found %v", 13, res)
	}

	res, _ = set.UpperBound(T{value: 19})
	if res == nil || (*res).value != 20 {
		t.Errorf("expected: %d, found %v", 20, res)
	}

	res, _ = set.UpperBound(T{value: 13})
	if res == nil || (*res).value != 14 {
		t.Errorf("expected: %d, found %v", 14, res)
	}

	res, _ = set.UpperBound(T{value: 23})
	if res != nil {
		t.Errorf("expected: nil, found %v", res)
	}
}

func TestUpperBoundStress(t *testing.T) {
	set := New[T]()
	l := 1
	r := 1000000
	for i := l; i <= r; i += 2 {
		set.Insert(T{value: i})
	}

	for i := l - 1; i < r; i += 2 {
		res, found := set.UpperBound(T{value: i})
		if found == false || (*res).value != i+1 {
			t.Errorf("expected: %d, found: %v", i+1, res)
		}
	}
}

func TestUpperBoundPair(t *testing.T) {
	set := New[Pair]()
	set.Insert(Pair{
		first:  1,
		second: 2,
	})
	set.Insert(Pair{
		first:  2,
		second: 3,
	})

	_, found := set.UpperBound(Pair{first: 1, second: 0})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should be found", 1, 0)
	}

	_, found = set.UpperBound(Pair{first: 1, second: 1})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should be found", 1, 1)
	}

	_, found = set.UpperBound(Pair{first: 2, second: 4})
	if found {
		t.Errorf("With pair: {first: %v, second: %v} should not be found", 2, 4)
	}

	_, found = set.UpperBound(Pair{first: 2, second: 2})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should not be found", 2, 2)
	}
}
