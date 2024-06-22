package main

import "fmt"

func main() {
	fmt.Println(reverseBits(1))

}

func reverseBits(num uint32) uint32 {
	var revNum uint32
	for range 32 {
		revNum = revNum*2 + num&1
		num >>= 1
	}
	return revNum
}
