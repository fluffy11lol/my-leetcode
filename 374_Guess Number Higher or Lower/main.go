package main

func main() {
	guessNumber(10)
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(n int) int {
	return 0
}
func guessNumber(n int) int {
	l := 0
	r := n
	m := (l + r) / 2
	var guessN = 1
	for l <= r {
		guessN = guess(m)
		if guessN == -1 {
			r = m - 1
		} else if guessN == 1 {
			l = m + 1
		} else {
			return m
		}
		m = (l + r) / 2
	}
	return 0
}
