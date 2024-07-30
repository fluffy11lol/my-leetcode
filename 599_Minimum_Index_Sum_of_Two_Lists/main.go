package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRestaurant([]string{"happy", "sad", "good"}, []string{"sad", "happy", "good"}))

}
func findRestaurant(list1 []string, list2 []string) []string {
	sMap := make(map[string]int)
	leastSum := math.MaxInt
	for i := range list1 {
		sMap[list1[i]] = i
	}
	res := make([]string, 0)
	for i := range list2 {
		if j, ok := sMap[list2[i]]; ok {
			if j+i < leastSum {
				leastSum = j + i
				res = []string{list2[i]}
			} else if j+i == leastSum {
				res = append(res, list2[i])
			}
		}
	}
	return res
}
