package main

func main() {

}
func bitwiseComplement(n int) int {
	n0 := n
	d := 1
	for n > 1 {
		d++
		n /= 2
	}
	return (1 << (d)) - 1 - n0
}
