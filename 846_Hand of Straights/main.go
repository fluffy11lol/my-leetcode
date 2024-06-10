package main

import "fmt"

func main() {
	fmt.Println(isNStraightHand([]int{1, 1, 2, 2, 3, 3}, 3))

}
func isNStraightHand(hand []int, groupSize int) bool {
	freqMap := make(map[int]int)
	if len(hand)%groupSize != 0 {
		return false
	}
	minN := hand[0]
	maxN := hand[0]
	for _, n := range hand {
		freqMap[n]++
		if n < minN {
			minN = n
		} else if n > maxN {
			maxN = n
		}
	}
	curN := minN
	for i := 0; i != len(hand); {
		for freqMap[curN] == 0 && curN < maxN {
			curN++
		}
		if freqMap[curN] == 0 {
			return false
		}
		c := freqMap[curN]
		freqMap[curN] = 0
		i += c * groupSize
		for j := curN + 1; j < curN+groupSize; j++ {
			if freqMap[j] < c {
				return false
			} else {
				freqMap[j] -= c
			}
		}
	}
	return true
}
