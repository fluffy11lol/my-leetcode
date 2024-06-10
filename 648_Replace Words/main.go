package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{
		"cat", "abat", "rat",
	}, ":ddd"))

}
func replaceWords(dictionary []string, sentence string) string {
	parts := strings.Split(sentence, " ")
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	for i, v := range parts {
		for _, vv := range dictionary {
			if strings.HasPrefix(v, vv) {
				parts[i] = vv
				break
			}
		}
	}
	result := strings.Join(parts, " ")
	return result
}
