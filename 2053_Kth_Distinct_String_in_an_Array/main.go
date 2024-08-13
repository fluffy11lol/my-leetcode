package main

func main() {

}
func kthDistinct(arr []string, k int) string {

	countMap := make(map[string]int)
	distMap := make(map[string]int)
	for i := range arr {
		countMap[arr[i]]++
	}
	c := 0
	for i := range arr {
		if countMap[arr[i]] == 1 {
			c++
			distMap[arr[i]] = c
			if c == k {
				return arr[i]
			}
		}
	}
	return ""
}
