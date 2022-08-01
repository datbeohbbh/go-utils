# Map Implementation

- Use Treap as the core algorithm

- You must implement the following methods on your own type so that you can use map implementation:

```golang
// Your Type
type K struct {
	value int
}

// Less comparison
func (t K) Less(rhs K) bool {
	return t.value < rhs.value
}

// Equal comparison
func (t K) Equal(rhs K) bool {
	return (t.value == rhs.value)
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