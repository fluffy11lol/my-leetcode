package main

import "fmt"

func main() {
	fmt.Println(gcd(4, 12))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	l := head
	curNode := head
	var gcdNode *ListNode
	for curNode.Next != nil {
		gcdNode = &ListNode{
			gcd(curNode.Val, curNode.Next.Val),
			curNode.Next,
		}
		curNode.Next = gcdNode
		curNode = gcdNode.Next
	}
	return l
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
