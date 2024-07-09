package main

func main() {

}
func averageWaitingTime(customers [][]int) float64 {
	sumWaitingTime := 0
	curTime := customers[0][0]
	for i := range customers {
		if customers[i][0] > curTime {
			curTime = customers[i][0]
		}
		curTime += customers[i][1]
		sumWaitingTime += curTime - customers[i][0]
	}
	return float64(sumWaitingTime) / float64(len(customers))
}
