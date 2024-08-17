package main

import "sort"

func main() {

}

func smallestDistancePair(nums []int, k int) int {

	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]-nums[0]

	for l < r {
		guess := (l + r) / 2
		count := 0
		i, j := 0, 1
		for i < len(nums) {
			for j < len(nums) && nums[j]-nums[i] <= guess {
				j++
			}
			count = count + (j - i - 1)
			i++
		}

		if count >= k {
			r = guess
		} else {
			l = guess + 1
		}
	}
	return l
}
