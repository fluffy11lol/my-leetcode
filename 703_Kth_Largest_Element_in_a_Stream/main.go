package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}

type KthLargest struct {
	k      int
	stream []int
}

func Constructor(k int, nums []int) KthLargest {
	copy(nums, nums)
	sort.Ints(nums)
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	insertPos, _ := slices.BinarySearch(this.stream, val)
	this.stream = append(this.stream[:insertPos], append([]int{val}, this.stream[insertPos:]...)...)
	return this.GetMaxK()
}
func (this *KthLargest) GetMaxK() int {
	return this.stream[len(this.stream)-this.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
