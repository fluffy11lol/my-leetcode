package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))

}
func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)

	projects := make([][2]int, n)
	for i := range n {
		projects[i] = [2]int{capital[i], profits[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	h := &MaxHeap{}
	heap.Init(h)

	for i, j := 0, 0; i < k; i++ {
		for ; j < n && projects[j][0] <= w; j++ {
			heap.Push(h, projects[j][1])
		}

		if h.Len() == 0 {
			break
		}

		w += heap.Pop(h).(int)
	}

	return w
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
