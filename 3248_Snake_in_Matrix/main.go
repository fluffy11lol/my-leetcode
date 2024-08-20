package main

func main() {

}
func finalPositionOfSnake(n int, commands []string) int {
	cur := 0
	moves := make(map[string]int)
	for i := range commands {
		moves[commands[i]]++
	}
	for command, v := range moves {
		switch command {
		case "RIGHT":
			cur += v
		case "LEFT":
			cur -= v
		case "UP":
			cur -= v * n
		case "DOWN":
			cur += v * n
		}
	}
	return cur
}
