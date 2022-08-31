package main

import (
	"bufio"
	"fmt"
	"os"

	li "github.com/datbeohbbh/go-utils/linked-list"
	"github.com/datbeohbbh/go-utils/numbers"
	"github.com/datbeohbbh/go-utils/queue"
)

type Int = numbers.Integer32

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

// bfs algorithm using linked list to save adjacent list.
func main() {
	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	adj := make([]*li.LinkedList[Int], N+1)
	for i := 1; i <= N; i++ {
		adj[i] = li.New[Int]()
	}

	for i := 1; i <= M; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u].Insert(numbers.NewInteger32(v))
		adj[v].Insert(numbers.NewInteger32(u))
	}

	// use queue package from github.com/datbeohbbh/go-utils/queue
	qu := queue.New[int32]()
	qu.Push(1)

	vis := make([]bool, N+1)
	vis[1] = true

	for !qu.Empty() {
		top := *qu.Front()
		qu.Pop()

		next := adj[top].Iterator()
		for {
			nxt := next()
			if nxt == nil {
				break
			}
			v := nxt.GetValue()
			if !vis[v] {
				fmt.Fprintf(writer, "%d -> %d\n", top, v)
				qu.Push(v)
				vis[v] = true
			}
		}
	}
}
