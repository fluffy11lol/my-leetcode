package main

import "fmt"

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))

}
func findArray(pref []int) []int {
	sourceArr := make([]int, len(pref))
	sourceArr[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		sourceArr[i] = pref[i-1] ^ pref[i]
	}
	return sourceArr
}
