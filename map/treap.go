package treap

import (
	"math"
	"math/rand"
	"time"

	compare "github.com/datbeohbbh/go-utils/map/interfaces"
)

const (
	LOWER_RANGE int64 = 1
	UPPER_RANGE int64 = math.MaxInt64
)

type node[K compare.IComparator[K], V any] struct {
	key      K
	value    V
	size     int64
	priority int64
	left     *node[K, V]
	right    *node[K, V]
}

func randPriority() int64 {
	rand.Seed(time.Now().UnixNano())
	return LOWER_RANGE + rand.Int63n(UPPER_RANGE)
}
