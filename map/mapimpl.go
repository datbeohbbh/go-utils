package treap

import compare "github.com/datbeohbbh/go-utils/interfaces"

type treeMap[K compare.IComparator[K], V any] struct {
	tree *node[K, V]
}

func New[K compare.IComparator[K], V any]() *treeMap[K, V] {
	return &treeMap[K, V]{
		tree: nil,
	}
}

func (curTreeMap *treeMap[K, V]) search(key K) (*node[K, V], bool) {
	res, existed := curTreeMap.tree.search(key)
	return res, existed
}

func (curTreeMap *treeMap[K, V]) Get(key K) *V {
	res, _ := curTreeMap.tree.search(key)
	if res == nil {
		return nil
	} else {
		return &(res.value)
	}
}

func (curTreeMap *treeMap[K, V]) Contains(key K) bool {
	_, existed := curTreeMap.search(key)
	return existed
}

func (curTreeMap *treeMap[K, V]) Put(key K, value V) {
	if curTreeMap.Contains(key) {
		curTreeMap.Erase(key)
	}
	curTreeMap.tree = curTreeMap.tree.put(&node[K, V]{
		key:      key,
		value:    value,
		size:     1,
		priority: randPriority(),
		left:     nil,
		right:    nil,
	})
}

func (curTreeMap *treeMap[K, V]) Erase(key K) bool {
	erased := true
	curTreeMap.tree, erased = curTreeMap.tree.erase(key)
	return erased
}

func (curTreeMap *treeMap[K, V]) Size() int64 {
	return curTreeMap.tree.size
}

func (curTreeMap *treeMap[K, V]) Empty() bool {
	return curTreeMap.tree.empty()
}

func (curTreeMap *treeMap[K, V]) Clear() {
	curTreeMap.tree = nil
}

func (curTreeMap *treeMap[K, V]) Begin() (*K, *V) {
	res := curTreeMap.tree.begin()
	if res == nil {
		return nil, nil
	} else {
		return &(res.key), &(res.value)
	}
}

func (curTreeMap *treeMap[K, V]) End() (*K, *V) {
	res := curTreeMap.tree.end()
	if res == nil {
		return nil, nil
	} else {
		return &(res.key), &(res.value)
	}
}

func (curTreeMap *treeMap[K, V]) LowerBound(key K) (*K, *V) {
	res := curTreeMap.tree.lowerBound(key)
	if res == nil {
		return nil, nil
	} else {
		return &(res.key), &(res.value)
	}
}

func (curTreeMap *treeMap[K, V]) UpperBound(key K) (*K, *V) {
	res := curTreeMap.tree.upperBound(key)
	if res == nil {
		return nil, nil
	} else {
		return &(res.key), &(res.value)
	}
}
