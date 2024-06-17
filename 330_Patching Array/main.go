package main

func main() {

}
func minPatches(nums []int, n int) int {
	miss := 1
	res := 0
	for i := 0; miss <= n; {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i += 1
		} else {
			miss += miss
			res++
		}
	}
	return res
}
