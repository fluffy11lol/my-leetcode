package main

func main() {

}
func maximumWealth(accounts [][]int) int {
	maxW := 0
	for i := range accounts {
		curW := 0
		for _, w := range accounts[i] {
			curW += w
		}
		if curW > maxW {
			maxW = curW
		}
	}
	return maxW
}
