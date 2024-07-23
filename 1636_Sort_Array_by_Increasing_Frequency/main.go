package main

import "sort"

func main() {

}
func frequencySort(nums []int) []int {
	mapp := make(map[int]int)
	n := len(nums)
	if n == 1 {
		return nums
	}
	for i := 0; i < n; i++ {
		mapp[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if mapp[nums[i]] == mapp[nums[j]] {
			return nums[i] > nums[j]
		}
		return mapp[nums[i]] < mapp[nums[j]]
	})
	return nums
}
