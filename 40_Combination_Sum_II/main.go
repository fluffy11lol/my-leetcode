package main

import (
	"sort"
)

func main() {

}
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	current := make([]int, 0)

	var backtrack func(int, int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > remain {
				break
			}

			current = append(current, candidates[i])
			backtrack(i+1, remain-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target)
	return result
}
