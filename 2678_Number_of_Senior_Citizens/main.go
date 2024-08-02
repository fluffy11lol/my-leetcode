package main

import "strconv"

func main() {

}
func countSeniors(details []string) int {
	c := 0
	for i := range details {
		if a, _ := strconv.Atoi(details[i][11:13]); a > 60 {
			c++
		}
	}
	return c
}
