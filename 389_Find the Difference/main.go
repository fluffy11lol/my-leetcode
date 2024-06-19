package main

func main() {

}
func findTheDifference(s string, t string) byte {
	res := 0
	for _, ch := range s {
		res -= int(ch)
	}
	for _, ch := range t {
		res += int(ch)
	}
	return byte(res)
}
