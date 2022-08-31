package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/datbeohbbh/go-utils/queue"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

// bfs algorithm
func main() {
	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	adj := make([][]int, N+1)

	for i := 1; i <= N; i++ {
		adj[i] = []int{}
	}

	for i := 1; i <= M; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// use queue package from github.com/datbeohbbh/go-utils/queue
	qu := queue.New[int]()
	qu.Push(1)

	vis := make([]bool, N+1)
	vis[1] = true

	for !qu.Empty() {
		top := *qu.Front()
		qu.Pop()

		for _, v := range adj[top] {
			if !vis[v] {
				fmt.Fprintf(writer, "%d -> %d\n", top, v)
				qu.Push(v)
				vis[v] = true
			}
		}
	}
}
