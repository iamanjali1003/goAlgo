package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	node, weight int
}

type MinHeap [][]int // [distance, node]

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func dijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt64
	}
	dist[start] = 0

	h := &MinHeap{{0, start}}
	heap.Init(h)

	for h.Len() > 0 {
		current := heap.Pop(h).([]int)
		currentDist, currentNode := current[0], current[1]

		if currentDist > dist[currentNode] {
			continue
		}

		for _, edge := range graph[currentNode] {
			distance := currentDist + edge.weight
			if distance < dist[edge.node] {
				dist[edge.node] = distance
				heap.Push(h, []int{distance, edge.node})
			}
		}
	}
	return dist
}

func main() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}
	start := 0
	distances := dijkstra(graph, start)
	fmt.Println("Shortest distances from node 0:", distances)
}
