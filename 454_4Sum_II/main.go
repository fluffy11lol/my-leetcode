package main

func main() {

}
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dMap := make(map[int]int, len(nums1)*len(nums2))
	for i := range nums1 {
		for j := range nums2 {
			dMap[nums1[i]+nums2[j]]++
		}
	}
	c := 0
	for i := range nums3 {
		for j := range nums4 {
			c += dMap[-nums3[i]-nums4[j]]
		}
	}
	return c
}
