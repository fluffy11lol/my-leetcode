package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sList := new(ListNode)
	resHead := sList
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			sList.Next = p1
			sList = sList.Next
			p1 = p1.Next
		} else {
			sList.Next = p2
			sList = sList.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		sList.Next = p1
	}
	if p2 != nil {
		sList.Next = p2
	}
	return resHead.Next
}
