package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(15, 4))

}
func numWaterBottles(numBottles int, numExchange int) int {
	n := 0
	usedBottles := 0
	for numBottles > 0 {
		n += numBottles
		usedBottles += numBottles
		numBottles = usedBottles / numExchange
		usedBottles = usedBottles % numExchange
	}
	return n
}
