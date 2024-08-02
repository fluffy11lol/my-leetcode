package main

func main() {

}
func minSwaps(nums []int) int {
	n := len(nums)
	totalOnes := 0

	// Count total number of 1's
	for _, num := range nums {
		totalOnes += num
	}

	// Edge cases
	if totalOnes == 0 || totalOnes == n {
		return 0
	}

	currentOnes := 0

	// Count 1's in the first window of size totalOnes
	for i := 0; i < totalOnes; i++ {
		currentOnes += nums[i]
	}

	maxOnes := currentOnes

	// Use two pointers to slide the window
	for i := 0; i < n; i++ {
		currentOnes -= nums[i]
		currentOnes += nums[(i+totalOnes)%n]
		if currentOnes > maxOnes {
			maxOnes = currentOnes
		}
	}

	return totalOnes - maxOnes
}
