package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))

}
func findCenter(edges [][]int) int {
	if edges[1][0] == edges[0][0] || edges[1][1] == edges[0][0] {
		return edges[0][0]
	}
	return edges[0][1]
}
