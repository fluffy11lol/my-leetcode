package main

import "strings"

func main() {

}
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseParentheses(s string) string {
	var stack []string
	for _, char := range s {
		if char == '(' {
			stack = append(stack, "(")
		} else if char == ')' {
			var rev strings.Builder
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				str := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				rev.WriteString(reverseString(str))
			}
			stack = stack[:len(stack)-1] // Pop the '('
			stack = append(stack, rev.String())
		} else {
			stack = append(stack, string(char))
		}
	}

	var result strings.Builder
	for _, str := range stack {
		result.WriteString(str)
	}

	return result.String()
}
