package main

func main() {

}
func getMinDistance(nums []int, target int, start int) int {
	d := 0
	for ; start-d >= 0 && start+d < len(nums); d++ {
		if nums[start-d] == target || nums[start+d] == target {
			return d
		}
	}
	for ; start-d >= 0; d++ {
		if nums[start-d] == target {
			return d
		}
	}
	for ; start+d < len(nums); d++ {
		if nums[start+d] == target {
			return d
		}
	}
	return -1
}
