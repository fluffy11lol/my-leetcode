package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))

}

// seq len = 10
func findRepeatedDnaSequences(s string) []string {
	dnaMap := make(map[string]int)
	for i := 0; i+10 <= len(s); i++ {
		dnaMap[s[i:i+10]]++
	}
	res := make([]string, 0)
	for k, v := range dnaMap {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
