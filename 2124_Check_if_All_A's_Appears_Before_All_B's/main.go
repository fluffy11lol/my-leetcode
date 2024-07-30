package main

func main() {

}
func checkString(s string) bool {
	wasB := false
	for i := range s {
		if s[i] == 'b' {
			wasB = true
		} else if wasB {
			return false
		}
	}
	return true
}
