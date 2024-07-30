package main

import (
	"sort"
	"strings"
)

func main() {

}
func sortSentence(s string) string {
	sArr := strings.Split(s, " ")
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i][len(sArr[i])-1] < sArr[j][len(sArr[j])-1]
	})
	for i := range sArr {
		sArr[i] = sArr[i][:len(sArr[i])-1]
	}
	return strings.Join(sArr, " ")
}
