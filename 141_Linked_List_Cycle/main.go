package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	lMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := lMap[head]; ok {
			return true
		} else {
			lMap[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}
