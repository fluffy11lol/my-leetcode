package main

import "fmt"

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))

}
func maximumGain(s string, x int, y int) int {
	var tar byte
	if x >= y {
		tar = 'a'
	} else {
		tar = 'b'
	}

	i := 0
	res := 0
	for i < len(s) {
		for i < len(s) && s[i] != 'a' && s[i] != 'b' {
			i++
		}
		a := 0
		b := 0
		cnt := 0
		pair := 0
		for i < len(s) && (s[i] == 'a' || s[i] == 'b') {
			if s[i] == 'a' {
				a++
			} else {
				b++
			}
			if s[i] == tar {
				cnt++
			} else {
				if cnt != 0 {
					cnt--
					pair++
				}
			}
			i++
		}
		a -= pair
		b -= pair
		res += pair*max(x, y) + min(a, b)*min(x, y)
	}
	return res
}
