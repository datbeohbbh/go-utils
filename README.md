## **Go Utils**

## **How to use?**

### Get from github

```bash
go get github.com/datbeohbbh/go-utils
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

## [*Qu - Queue*](https://github.com/datbeohbbh/go-utils/tree/master/queue)

#### [Queue example](https://github.com/datbeohbbh/go-utils/examples/queue/main.go)

### Types

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/queue/node.go)
```go
type node[T any] struct {
	next *node[T]
	value T
}
```

#### [type queue](https://github.com/datbeohbbh/go-utils/blob/master/queue/queueImpl.go)
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


#### func [Back](https://github.com/datbeohbbh/go-utils/blob/master/queue/back.go)
```go
func (qu *Queue[T]) Back() *T
```
returns a pointer to the last element in queue or `nil` is queue is empty.

#### func [Push](https://github.com/datbeohbbh/go-utils/blob/master/queue/push.go)
```go
func (qu *Queue[T]) Push(value T)
```
pushes a new element into the queue.

#### func [Pop](https://github.com/datbeohbbh/go-utils/blob/master/queue/pop.go)
```go
func (qu *Queue[T]) Pop()
```
remove the first element in the queue.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/queue/size.go)
```go
func (qu *Queue[T]) Size() int64
```
returns size of the queue.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/queue/empty.go)
```go
func (qu *Queue[T]) Empty() bool
```
returns `true` if queue is empty, otherwise `false`.

#### func [Clear](https://github.com/datbeohbbh/go-utils/blob/master/queue/clear.go)
```go
func (qu *Queue[T]) Clear()
```
clears current queue.


## [li - Linked List](https://github.com/datbeohbbh/go-utils/tree/master/linked-list)

#### [Linked list example](https://github.com/datbeohbbh/go-utils/examples/linkedlist/main.go)

### Types

#### [type node](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/node.go)
```go
type node[T compare.IEqual[T]] struct {
	value T
	next *node[T]
}
```
type parameter `T` must implement `Equal(T) bool` method of interface [IEqual](https://github.com/datbeohbbh/go-utils/blob/master/interfaces/comparator.go)

#### [type linkedList](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/linkedListImpl.go)
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
inserts a new value into the linked list.

#### func [Remove](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/remove.go)
```go
func (li *LinkedList[T]) Remove(value T) bool
```
removes the first element equal to `value` in linked list. Returns `true` if there is an element equal to `value` in linked list, otherwise `false`.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/empty.go)
```go
func (li *LinkedList[T]) Empty() bool
```
returns `true` if linked list is empty, otherwise `false`.

#### func [GetHead](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/getHead.go)
```go
func (li *LinkedList[T]) GetHead() *T
```
returns a pointer to the first element in linked list. Returns `nil` if linked list is empty.

#### func [GetTail](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/getTail.go)
```go
func (li *LinkedList[T]) GetTail() *T
```
returns a pointer to the last element in linked list. Returns `nil` if linked list is empty.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/size.go)
```go
func (li *LinkedList[T]) Size() int64
```
returns size of the linked list.

#### func [Search](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/search.go)
```go
func (li *LinkedList[T]) Search(value T) bool
```
returns `true` if there is an element equal to `value` in linked list, otherwise `false`.

#### func [Iterator](https://github.com/datbeohbbh/go-utils/blob/master/linked-list/iterator.go)
```go
func (li *LinkedList[T]) Iterator() func() *T
```
returns the lazy function, so that each time that function is invoked, it returns `next` element in the linked list or `nil` if reached the end of linked list (check example for more details).