package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("-91283472332"))

}
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	p := 0
	if strings.ContainsAny("1234567890+-", string(s[0])) {
		p++
	} else {
		return 0
	}
	for _, c := range s[1:] {
		if strings.ContainsAny("1234567890", string(c)) {
			p++
			continue
		} else {
			break
		}
	}
	i, _ := strconv.Atoi(s[:p])
	if i > math.MaxInt32 {
		return math.MaxInt32
	}
	if i < math.MinInt32 {
		return math.MinInt32
	}
	return i
}
