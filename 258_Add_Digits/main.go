package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))

}
func addDigits(num int) int {
	temp := 0
	for num > 9 {
		temp = 0
		for num > 0 {
			temp += num % 10
			num /= 10
		}
		num = temp
	}
	return num
}
