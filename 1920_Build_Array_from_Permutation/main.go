package main

func main() {

}
func buildArray(nums []int) []int {
	for i, n := range nums {
		nums[i] += (nums[n] % 1000) * 1000
	}
	for i := range nums {
		nums[i] /= 1000
	}
	return nums
}
