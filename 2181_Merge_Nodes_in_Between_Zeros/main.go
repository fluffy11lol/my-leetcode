package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	sumListSentinel := new(ListNode)
	curSumNode := sumListSentinel
	sumListSentinel.Next = curSumNode
	curNode := head.Next
	curSum := 0
	for curNode != nil {
		if curNode.Val == 0 {
			curSumNode.Next = new(ListNode)
			curSumNode = curSumNode.Next
			curSumNode.Val = curSum
			curSum = 0
		} else {
			curSum += curNode.Val
		}
		curNode = curNode.Next
	}
	return sumListSentinel.Next
}
