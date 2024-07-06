package main

import (
	"fmt"
	"math"
)

func main() {
	head := new(ListNode)
	curNode := new(ListNode)
	curNode.Val = 5
	head.Next = curNode
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 3
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 1
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 2
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 5
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 1
	curNode.Next = new(ListNode)
	curNode = curNode.Next
	curNode.Val = 2
	fmt.Println(nodesBetweenCriticalPoints(head.Next))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	maxD, minD := -1, math.MaxInt
	sI, bI, curI := -1, -1, 0
	prevN, nextN := 0, 0
	curN := head.Val
	head = head.Next

	for head.Next != nil {
		prevN = curN
		nextN = head.Next.Val
		curN = head.Val
		if curN > prevN && curN > nextN || curN < prevN && curN < nextN {
			if sI == -1 {
				sI = curI
			} else {
				maxD = curI - sI
				minD = min(minD, curI-bI)
			}
			bI = curI
		}
		curI++
		head = head.Next
	}

	if sI != bI {
		return []int{minD, maxD}
	} else {
		return []int{-1, -1}
	}
}
