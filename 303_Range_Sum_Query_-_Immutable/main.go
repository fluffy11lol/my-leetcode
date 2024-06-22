package main

func main() {

}

type NumArray struct {
	nums      []int
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{
		nums:      nums,
		prefixSum: make([]int, len(nums)),
	}
	pSum := 0
	for i, n := range nums {
		pSum += n
		arr.prefixSum[i] = pSum
	}
	return arr
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left] + this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
