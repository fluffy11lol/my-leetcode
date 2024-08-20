package main

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))

}

func brokenCalc(startValue int, target int) int {
	steps := 0
	for startValue < target {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		steps++
	}
	return steps + startValue - target
}
