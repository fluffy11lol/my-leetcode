package main

func main() {

}
func minSum(nums1 []int, nums2 []int) int64 {
	curS1, curS2 := 0, 0
	z1, z2 := 0, 0
	for i := range nums1 {
		if nums1[i] == 0 {
			z1++
		}
		curS1 += nums1[i]
	}
	for i := range nums2 {
		if nums2[i] == 0 {
			z2++
		}
		curS2 += nums2[i]
	}
	minS1, minS2 := curS1+z1, curS2+z2
	if z1 == 0 && minS2 > curS1 || z2 == 0 && minS1 > curS2 {
		return -1
	}
	return int64(max(minS1, minS2))
}
