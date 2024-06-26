package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"reflower", "flow", "flight"}))

}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minWord := strs[0]
	for _, word := range strs[1:] {
		if len(word) < len(minWord) {
			minWord = word
		}
	}

	curPrefix := minWord
	curLength := len(minWord)
OuterLoop:
	for curLength > 0 {
		for _, word := range strs {
			if !(word[:curLength] == curPrefix) {
				curLength--
				curPrefix = curPrefix[:len(curPrefix)-1]
				continue OuterLoop
			}
		}
		return curPrefix
	}
	return ""
}
