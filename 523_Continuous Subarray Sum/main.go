package main

func main() {

}
func checkSubarraySum(nums []int, k int) bool {
	// Map to store the remainders and their corresponding index
	remaindersFound := make(map[int]int)
	currSum := 0
	remaindersFound[0] = -1 // To handle the case when subarray starts from index 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		remainder := currSum % k

		if val, exists := remaindersFound[remainder]; exists {
			// Check if the length of the subarray is at least 2
			if i-val >= 2 {
				return true
			}
		} else {
			// Store the remainder and the current index
			remaindersFound[remainder] = i
		}
	}

	return false
}
