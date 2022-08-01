# Map Implementation

- Use Treap as the core algorithm

- You must implement the following methods on your own type so that you can use map implementation:

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
	treeMap "github.com/datbeohbbh/go-utils/map"
)

func main() {
	treeMap := treeMap.New[K, int]()	
}
```