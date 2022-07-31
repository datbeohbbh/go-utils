# Set Implementation

- Use AVL Tree as the core algorithm

- You must implement the following methods on your own type so that you can use set implementation:

```golang
type T struct {
	value int
}

// Type assert to cast from type any to your own type
func (t T) Cast(rhs any) T {
	x, same := rhs.(T)
	if !same {
		log.Panic("type error!!")
	}
	return x
}

// Less comparison
func (t T) Less(rhs any) bool {
	return t.value < t.Cast(rhs).value
}

// Equal comparison
func (t T) Equal(rhs any) bool {
	return (t.value == t.Cast(rhs).value)
}
```

- How to use?

```golang
import (
	set "github.com/datbeohbbh/go-utils/set"
)

func main() {
	s := set.New[T]()
	s.Insert(T{value: 10})
	s.Remove(T{value: 10})
	found := s.Find(T{value: 10})
}
```