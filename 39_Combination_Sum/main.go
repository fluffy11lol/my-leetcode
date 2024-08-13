package main

import "slices"

func main() {

}
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	ans := make([][]int, 0)
	backTrack(&ans, []int{}, candidates, target, 0)
	return ans
}
func backTrack(ans *[][]int, combination []int, candidates []int, target int, startIndex int) {
	if target == 0 {
		validComb := make([]int, len(combination))
		copy(validComb, combination)
		*ans = append(*ans, validComb)
		return
	}
	if target < 0 {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		backTrack(ans, append(combination, candidates[i]), candidates, target-candidates[i], i)
	}
}
