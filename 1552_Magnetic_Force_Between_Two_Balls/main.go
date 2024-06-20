package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{94885264, 34193493, 976874764, 752959504, 746503881, 814985479, 691787337, 270416883, 603014310, 602556426, 644963448, 333443157, 70956635, 857223019, 249457690, 782919344, 171062134, 469997994, 792308825, 289080527, 100536829, 764630760, 174433254, 431195146, 178921163, 405666805, 772490742, 126675275, 723294793, 56238254, 17964069, 659279343, 451832985, 448396998, 36553750, 200797149, 762601607, 297120983, 756162959, 459654861}, 15))

}
func maxDistance(position []int, m int) int {

	sort.Ints(position)
	l, r := 1, position[len(position)-1]-position[0]
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		lastPosition, balls := position[0], 1
		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= mid {
				lastPosition = position[i]
				balls++
			}
		}
		if balls >= m {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
