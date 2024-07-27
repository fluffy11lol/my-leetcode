package main

func main() {

}
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	n := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			n++
		}
	}
	return n
}
