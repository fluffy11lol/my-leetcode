package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(701))

}
func convertToTitle(columnNumber int) string {
	resStr := ""
	for columnNumber > 0 {
		columnNumber--
		rem := columnNumber % 26
		resStr = string(rune('A'+rem)) + resStr
		columnNumber /= 26
	}
	return resStr
}
