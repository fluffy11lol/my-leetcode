package main

func main() {

}
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	locMaxes := make([][]int, n-2)
	for i := range locMaxes {
		locMaxes[i] = make([]int, n-2)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			maxVal := grid[i-1][j-1]
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if grid[x][y] > maxVal {
						maxVal = grid[x][y]
					}
				}
			}
			locMaxes[i-1][j-1] = maxVal
		}
	}

	return locMaxes
}
