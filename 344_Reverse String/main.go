package main

func main() {
	reverseString([]byte{1, 2, 3, 4, 5})
}
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
