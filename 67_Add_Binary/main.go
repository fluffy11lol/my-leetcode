package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	const one, zero = byte('1'), byte('0')
	binarySum := ""
	carry := "0"
	p1, p2 := len(a)-1, len(b)-1
	for p1 >= 0 {
		if a[p1] == b[p2] {
			if carry == "0" {
				binarySum = "0" + binarySum
			} else {
				binarySum = "1" + binarySum
			}
			carry = string(a[p1])
		} else if carry == "0" {
			binarySum = "1" + binarySum
		} else {
			binarySum = "0" + binarySum
			carry = "1"
		}
		p1--
		p2--
	}
	for p2 >= 0 {
		if b[p2] == one {
			if carry == "0" {
				binarySum = "1" + binarySum
			} else {
				binarySum = "0" + binarySum
				carry = "1"
			}
		} else {
			binarySum = carry + binarySum
			carry = "0"
		}
		p2--
	}
	if carry == "1" {
		binarySum = "1" + binarySum
	}
	return binarySum
}
