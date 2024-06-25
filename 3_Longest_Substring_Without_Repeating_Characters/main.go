package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

}
func lengthOfLongestSubstring(s string) int {
	ans := 0
	count := make([]int, 256)
	for i, j := 0, 0; i < len(s); i++ {
		count[s[i]]++
		for count[s[i]] > 1 {
			count[s[j]]--
			j++
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	return ans
}
