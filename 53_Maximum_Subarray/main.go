package main

func main() {

}
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curSum = max(curSum+num, num)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
