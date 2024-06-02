package main

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))

}
func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1
	m := (l + r) / 2
	for l <= r {
		if letters[m] > target {
			r = m - 1
		} else if letters[m] < target {
			l = m + 1
		} else {
			for letters[m] == target && m < (len(letters)-1) {
				m++
			}
			if letters[m] != target {
				return letters[m]
			} else {
				return letters[0]
			}
		}
		m = (l + r) / 2
	}
	if letters[m] > target {
		return letters[m]
	} else if m+1 < len(letters) {
		if letters[m+1] > target {
			return letters[m+1]
		}
	}
	return letters[0]
}
