package main

func main() {

}
func minimumOperations(nums []int) int {
	op := 0
	for _, n := range nums {
		if n%3 != 0 {
			op++
		}
	}
	return op
}
