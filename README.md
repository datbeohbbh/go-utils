## **Go Utils**

## **How to use?**

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

### [*Queue*](https://github.com/datbeohbbh/go-utils/tree/master/queue)

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
type queue[T any] struct {
	head *node[T]
	tail *node[T]	
	size int64
}
```

#### func [New](https://github.com/datbeohbbh/go-utils/blob/master/queue/queueImpl.go)
```go
func  New[T any]() *queue[T]
```
returns a pointer to the new queue.

#### func [Front](https://github.com/datbeohbbh/go-utils/blob/master/queue/front.go)
```go
func (qu *queue[T]) Front() *T
```
returns a pointer to the first element in queue or `nil` is queue is empty.


#### func [Back](https://github.com/datbeohbbh/go-utils/blob/master/queue/back.go)
```go
func (qu *queue[T]) Back() *T
```
returns a pointer to the last element in queue or `nil` is queue is empty.

#### func [Push](https://github.com/datbeohbbh/go-utils/blob/master/queue/push.go)
```go
func (qu *queue[T]) Push(value T)
```
pushes a new element into the queue.

#### func [Pop](https://github.com/datbeohbbh/go-utils/blob/master/queue/pop.go)
```go
func (qu *queue[T]) Pop()
```
remove the first element in the queue.

#### func [Size](https://github.com/datbeohbbh/go-utils/blob/master/queue/size.go)
```go
func (qu *queue[T]) Size() int64
```
returns size of the queue.

#### func [Empty](https://github.com/datbeohbbh/go-utils/blob/master/queue/empty.go)
```go
func (qu *queue[T]) Empty() bool
```
returns `true` if queue is empty, otherwise `false`.

#### func [Clear](https://github.com/datbeohbbh/go-utils/blob/master/queue/clear.go)
```go
func (qu *queue[T]) Clear()
```
clears current queue.