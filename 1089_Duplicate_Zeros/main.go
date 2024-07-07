package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)

}
func duplicateZeros(arr []int) {
	copyArr := make([]int, len(arr))
	j := 0
	//t := 0
	for i := 0; j < len(arr); i++ {
		if arr[i] == 0 {
			j++
		} else {
			copyArr[j] = arr[i]
		}
		j++
	}
	copy(arr, copyArr)
}
