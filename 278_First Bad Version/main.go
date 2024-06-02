package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return false
}
func firstBadVersion(n int) int {
	firstBad := 0
	l := 0
	r := n
	m := (l + r) / 2
	var test bool
	for l <= r {
		test = isBadVersion(m)
		if test {
			firstBad = m
			r = m - 1
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}
	return firstBad
}
