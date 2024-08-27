package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}
func fractionAddition(expression string) string {
	common, num := 5*7*8*9, 0
	for _, exps := range strings.Split(expression, "+") {
		for i, exp := range strings.Split(exps, "-") {
			if len(exp) == 0 {
				continue
			}
			fraction := strings.Split(exp, "/")
			u, _ := strconv.Atoi(fraction[0])
			d, _ := strconv.Atoi(fraction[1])
			if i == 0 {
				num += u * (common / d)
			} else {
				num -= u * (common / d)
			}
		}
	}
	for i := 2; i <= 7; i++ {
		for num%i == 0 && common%i == 0 {
			num /= i
			common /= i
		}
	}

	return fmt.Sprintf("%d/%d", num, common)
}
