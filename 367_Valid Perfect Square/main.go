package main

func main() {

}
func isPerfectSquare(num int) bool {
	l := 0
	r := num + 1
	mid := (l + r) / 2
	sqrt := 0
	for l <= r {
		sqrt = mid * mid
		if sqrt > num {
			r = mid - 1
		} else if sqrt < num {
			l = mid + 1
		} else {
			return true
		}
		mid = (l + r) / 2
	}
	return false
}
