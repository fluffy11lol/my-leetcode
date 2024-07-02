package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 3, 3, 3}, []int{1, 3, 4, 3}))

}
func intersect(nums1 []int, nums2 []int) []int {
	arr1Map := make(map[int]int)
	for _, num := range nums1 {
		arr1Map[num]++
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		if arr1Map[num] > 0 {
			res = append(res, num)
			arr1Map[num]--
		}
	}
	return res
}
