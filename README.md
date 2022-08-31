## **Go Utils**
*This data structure library is my small personal project in the summer of 2022. Its main purpose is that I use some of them for my bigger one.*

## **How to use?**

### Get from github

```bash
go get -u github.com/datbeohbbh/go-utils
```

### Play with source

 - Build docker image:
```bash
docker build -t utils/go-utils .
```

-	Run test:
```bash
docker run --rm utils/go-utils
```

- Work in docker container:
```bash
docker run -it --rm -v $(pwd):/go-utils utils/go-utils sh
```

## **Docs**

[Queue Doc]::

## [*Qu - Queue*](https://github.com/datbeohbbh/go-utils/tree/master/queue)

#### [Queue example](https://github.com/datbeohbbh/go-utils/blob/master/examples/queue/main.go)

### Types

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/queue/node.go)
```go
type node[T any] struct {
	next *node[T]
	value T
}
```

#### [type Queue](https://github.com/datbeohbbh/go-utils/blob/master/queue/queueImpl.go)
```go
type Queue[T any] struct {
	head *node[T]
	tail *node[T]	
	size int64
}
```

#### func [New](https://github.com/datbeohbbh/go-utils/blob/master/queue/queueImpl.go)
```go
func  New[T any]() *Queue[T]
```
returns a pointer to the new queue.

#### func [Front](https://github.com/datbeohbbh/go-utils/blob/master/queue/front.go)
```go
func (qu *Queue[T]) Front() *T
```
returns a pointer to the first element in queue or `nil` is queue is empty.
Complexity: *O(1)*.

#### func [Back](https://github.com/datbeohbbh/go-utils/blob/master/queue/back.go)
```go
func (qu *Queue[T]) Back() *T
```
returns a pointer to the last element in queue or `nil` is queue is empty.
Complexity: *O(1)*.

#### func [Push](https://github.com/datbeohbbh/go-utils/blob/master/queue/push.go)
```go
func (qu *Queue[T]) Push(value T)
```
pushes a new element into the queue.
Complexity: *O(1)*.

#### func [Pop](https://github.com/datbeohbbh/go-utils/blob/master/queue/pop.go)
```go
func (qu *Queue[T]) Pop()
```
remove the first element in the queue.
Complexity: *O(1)*.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/queue/size.go)
```go
func (qu *Queue[T]) Size() int64
```
returns size of the queue.
Complexity: *O(1)*.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/queue/empty.go)
```go
func (qu *Queue[T]) Empty() bool
```
returns `true` if queue is empty, otherwise `false`.
Complexity: *O(1)*.

#### func [Clear](https://github.com/datbeohbbh/go-utils/blob/master/queue/clear.go)
```go
func (qu *Queue[T]) Clear()
```
clears current queue.
Complexity: *O(1)*.

- - -
[Linked List Doc]::

## [li - Linked List](https://github.com/datbeohbbh/go-utils/tree/master/linked-list)

#### [Linked list example](https://github.com/datbeohbbh/go-utils/blob/master/examples/linkedlist/main.go)

### Types

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/node.go)
```go
type node[T compare.IEqual[T]] struct {
	value T
	next *node[T]
}
```
type parameter `T` must implement `Equal(T) bool` method of interface [IEqual](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go)

#### [type LinkedList](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/linkedListImpl.go)
```go
type LinkedList[T compare.IEqual[T]] struct {
	head *node[T]
	tail *node[T]
	size int64
}
```
type parameter `T` must implement `Equal(T) bool` method of interface [IEqual](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go)

#### func [New](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/linkedListImpl.go)
```go
func New[T compare.IEqual[T]]() *LinkedList[T]
```
- type parameter `T` must implement `Equal(T) bool` method of interface [IEqual](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).
- returns a pointer to the new linked list.

#### func [Insert](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/insert.go)
```go
func (li *LinkedList[T]) Insert(value T)
```
inserts a new value in the end of linked list.
Complexity: *O(1)*.

#### func [Remove](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/remove.go)
```go
func (li *LinkedList[T]) Remove(value T) bool
```
removes the first element equal to `value` in linked list. Returns `true` if there is an element equal to `value` in linked list, otherwise `false`.
Complexity: *O(SIZE_OF_LINKED_LIST)*.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/empty.go)
```go
func (li *LinkedList[T]) Empty() bool
```
returns `true` if linked list is empty, otherwise `false`.
Complexity: *O(1)*.

#### func [GetHead](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/getHead.go)
```go
func (li *LinkedList[T]) GetHead() *T
```
returns a pointer to the first element in linked list. Returns `nil` if linked list is empty.
Complexity: *O(1)*.

#### func [GetTail](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/getTail.go)
```go
func (li *LinkedList[T]) GetTail() *T
```
returns a pointer to the last element in linked list. Returns `nil` if linked list is empty.
Complexity: *O(1)*.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/size.go)
```go
func (li *LinkedList[T]) Size() int64
```
returns size of the linked list.
Complexity: *O(1)*.

#### func [Search](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/search.go)
```go
func (li *LinkedList[T]) Search(value T) bool
```
returns `true` if there is an element equal to `value` in linked list, otherwise `false`.
Complexity: *O(SIZE_OF_LINKED_LIST)*.

#### func [Iterator](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/iterator.go)
```go
func (li *LinkedList[T]) Iterator() func() *T
```
returns the lazy function, so that each time that function is invoked, it returns `next` element in the linked list or `nil` if reached the end of linked list (check example for more details).
Complexity: *O(SIZE_OF_LINKED_LIST)*.

- - -
[Map Doc]::

## [*map - TreeMap*](https://github.com/datbeohbbh/go-utils/tree/master/map)

#### [Map example](https://github.com/datbeohbbh/go-utils/blob/master/examples/map/main.go)

#### *`Map` is a sorted associative container that contains key-value pairs with unique keys. Here I used [Treap](https://cp-algorithms.com/data_structures/treap.html) as core algorithm.*

### Types 

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/map/treap.go) (not exported)
```go
type node[K compare.IComparator[K], V any] struct
```
Type parameter for key `K` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).

#### [type TreeMap](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
type TreeMap[K compare.IComparator[K], V any] struct
```
Type parameter for key `K` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).

#### func [New](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func New[K compare.IComparator[K], V any]() *TreeMap[K, V]
```
- Type parameter for key `K` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).
- return a pointer to the new tree map.

#### func [Put](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Put(key K, value V)
```
puts new key, value pair into the tree map. If key is already represented, old value is replaced value with `value`. 
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Erase](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Erase(key K) bool
```
erases node in tree map with key is `key`. Returns `true` if key `key` exists, otherwise `false`.
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Get](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Get(key K) *V
```
return a pointer to value of the node with key is `key`. Returns `nil` if key `key` is not represented in tree map.
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Contains](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Contains(key K) bool
```
returns `true` if key `key` is represented in tree map, otherwise `false`. 
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Size() int64
```
returns size of tree map. 
Complexity: *O(1)*.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Empty() bool
```
returns `true` if tree map is empty, otherwise `false`. 
Complexity: *O(1)*.

#### func [Begin](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Begin() (*K, *V)
```
returns the key (pointer), value (pointer) pair of the node with smallest key value, return `key = nil`, `value = nil` if tree map is empty. 
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Last](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Last() (*K, *V)
```
returns the key (pointer), value (pointer) pair of the node with biggest key value, return `key = nil`, `value = nil` if tree map is empty. 
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [LowerBound](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) LowerBound(key K) (*K, *V)
```
returns the key (pointer), value (pointer) pair of the node with smallest (according to your comparison method) key value greater than or equal `key`. Complexity: *O(log(SIZE_OF_MAP))*.

#### func [UpperBound](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) UpperBound(key K) (*K, *V)
```
returns the key (pointer), value (pointer) pair of the node with smallest (according to your comparison method) key value greater than `key`. 
Complexity: *O(log(SIZE_OF_MAP))*.

#### func [Iterator](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Iterator() func() (*K, *V)
```
returns the lazy function, so that each time that function is invoked, it returns the key (pointer), value (pointer) pair of next node with key in ascending order (check example for more details). 
Complexity (iterate over the tree map): *O(SIZE_OF_MAP * log(SIZE_OF_MAP))*

#### func [Clear](https://github.com/datbeohbbh/go-utils/blob/master/map/mapImpl.go)
```go
func (curTreeMap *TreeMap[K, V]) Clear()
```
clears the tree map.
Complexity: *O(1)*.

- - -
[Set Doc]::

## [*set - TreeSet*](https://github.com/datbeohbbh/go-utils/tree/master/set)

#### [Set example](https://github.com/datbeohbbh/go-utils/blob/master/examples/set/main.go)

#### *`Set` is an associative container that contains a sorted set of unique objects. Here I used [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree) as core algorithm.*

### Types 

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/set/avlTree.go) (not exported)
```go
type node[T compare.IComparator[T]] struct
```
Type parameter `T` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).

#### [type TreeSet](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
type TreeSet[T compare.IComparator[T]] struct
```
Type parameter `T` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).

#### func [New](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func New[T compare.IComparator[T]]() *TreeSet[T]
```
- Type parameter `T` must implement methods `Less(T) bool`, `Equal(T) bool` of interface [IComparator[T  any]](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go).
- returns a pointer to the new tree set.

#### func [Insert](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Insert(value T) bool
```
if `value` is not represented in set, `value` is inserted and returns true. Otherwise, return `false`.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [Remove](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Remove(value T) bool
```
if `value` is represented in set, `value` is removed and returns `true`. Otherwise, return `false`.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [Find](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Find(value T) bool
```
returns `true` is `value` is represented in set, otherwise `false`.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [Begin](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Begin() *T
```
returns a pointer to the value of the smallest (according to your comparison method) element in set. Returns `nil` if set is empty.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [Last](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Last() *T
```
returns a pointer to the value of the biggest (according to your comparison method) element in set. Returns `nil` if set is empty.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [LowerBound](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) LowerBound(value T) (*T, bool)
```
returns a pointer the the value of the smallest (according to your comparison method) element in set that greater than or equal `value`. Returns `nil` if there is no such element.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [UpperBound](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) UpperBound(value T) (*T, bool)
```
returns a pointer the the value of the smallest (according to your comparison method) element in set that greater than `value`. Returns `nil` if there is no such element.
Complexity: *O(log(SIZE_OF_SET))*.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Empty() bool
```
return `true` if set is empty, otherwise `false`.
Complexity: *O(1)*.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Size() int64
```
return the size of the set.
Complexity: *O(1)*.

#### func [Iterator](https://github.com/datbeohbbh/go-utils/blob/master/set/setImpl.go)
```go
func (curTreeSet *TreeSet[T]) Iterator() func() *T
```
returns the lazy function, so that each time that function is invoked, it returns the value (pointer) of next node with value in ascending order (check example for more details). 
Complexity (iterate over the set): *O(SIZE_OF_SET * log(SIZE_OF_SET))*.
