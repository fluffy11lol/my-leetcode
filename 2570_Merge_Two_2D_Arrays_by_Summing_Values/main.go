package main

func main() {

}
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	resArr := make([][]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] < nums2[j][0] {
			resArr = append(resArr, []int{nums1[i][0], nums1[i][1]})
			i++
		} else if nums1[i][0] > nums2[j][0] {
			resArr = append(resArr, []int{nums2[j][0], nums2[j][1]})
			j++
		} else {
			resArr = append(resArr, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		}
	}
	for ; i < len(nums1); i++ {
		resArr = append(resArr, []int{nums1[i][0], nums1[i][1]})
	}
	for ; j < len(nums2); j++ {
		resArr = append(resArr, []int{nums2[j][0], nums2[j][1]})
	}

	return resArr
}
