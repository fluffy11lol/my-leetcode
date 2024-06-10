package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))

}
func commonChars(words []string) []string {
	commonArr := make([]int, 26)
	for _, ch := range words[0] {
		commonArr[ch-'a']++
	}
	for i := 1; i < len(words); i++ {
		curArr := make([]int, 26)
		for _, ch := range words[i] {
			curArr[ch-'a']++
		}
		for r, c := range commonArr {
			if c != 0 && curArr[r] != 0 {
				if curArr[r] < c {
					commonArr[r] = curArr[r]
				}
			} else {
				commonArr[r] = 0
			}
		}
	}
	resArr := make([]string, 0)
	for i, c := range commonArr {
		if c != 0 {
			for range c {
				resArr = append(resArr, string(rune(i+'a')))
			}
		}
	}
	return resArr
}
