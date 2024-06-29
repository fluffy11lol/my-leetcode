package main

func main() {

}
func containsDuplicate(nums []int) bool {
	freqMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if freqMap[num] {
			return true
		}
		freqMap[num] = true
	}
	return false
}
