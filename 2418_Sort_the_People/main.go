package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))

}
func sortPeople(names []string, heights []int) []string {
	sortArr := make([]struct {
		name   string
		height int
	}, len(names))
	for i := range names {
		sortArr[i].name = names[i]
		sortArr[i].height = heights[i]
	}
	sort.Slice(sortArr, func(i, j int) bool {
		return sortArr[i].height > sortArr[j].height
	})
	resNames := make([]string, len(names))
	for i := range sortArr {
		resNames[i] = sortArr[i].name
	}
	return resNames
}
