# Set Implementation

- Use AVL Tree as the core algorithm

- You must implement the following methods on your own type so that you can use set implementation:

```golang

// Your type
type T struct {
	value int
}

// Less comparison
func (t T) Less(rhs T) bool {
	return t.value < rhs.value
}

// Equal comparison
func (t T) Equal(rhs T) bool {
	return (t.value == rhs.value)
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