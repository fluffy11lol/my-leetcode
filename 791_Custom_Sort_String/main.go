package main

import "fmt"

func main() {
	fmt.Println(customSortString("cba", "abcd"))

}
func customSortString(order string, s string) string {
	cMap := make(map[byte]int)
	resS := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		cMap[s[i]]++
	}
	n := 0
	for _, c := range order {
		for range cMap[byte(c)] {
			resS[n] = byte(c)
			n++
		}
		delete(cMap, byte(c))
	}
	for c := range cMap {
		for range cMap[c] {
			resS[n] = c
			n++
		}
	}
	return string(resS)
}
