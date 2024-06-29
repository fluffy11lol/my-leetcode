package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

}
func maxProfit(prices []int) int {
	minPrice := prices[0]
	res := 0
	for i := range prices {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}
	return res
}
