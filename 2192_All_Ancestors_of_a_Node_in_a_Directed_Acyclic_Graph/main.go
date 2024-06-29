package main

import "fmt"

func main() {
	n := 8
	edges := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
	ancestors := getAncestors(n, edges)
	fmt.Println(ancestors)
}

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	res := make([][]int, n)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	for i := 0; i < n; i++ {
		bfs(i, graph, res)
	}
	return res
}
func bfs(n int, graph [][]int, res [][]int) {
	queue := make([]int, 0)
	seen := make(map[int]bool, 0)
	seen[n] = true

	queue = append(queue, n)

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		for _, v := range graph[el] {
			if !seen[v] {
				seen[v] = true
				queue = append(queue, v)
				res[v] = append(res[v], n)
			}
		}
	}
}
