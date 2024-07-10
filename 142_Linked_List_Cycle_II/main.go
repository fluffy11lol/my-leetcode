package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	lMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := lMap[head]; ok {
			return head
		} else {
			lMap[head] = struct{}{}
		}
		head = head.Next
	}
	return nil
}
