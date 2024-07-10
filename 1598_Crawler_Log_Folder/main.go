package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))

}
func minOperations(logs []string) int {
	d := 0 // start in main
	for i := range logs {
		switch logs[i] {
		case "./":
			continue
		case "../":
			if d > 0 {
				d--
			}
		default:
			d++
		}
	}
	return d
}
