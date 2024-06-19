package main

func main() {
	l1 := new(ListNode)
	curNode := l1
	for i := range 4 {
		curNode.Val = i
		if i == 3 {
			break
		}
		curNode.Next = new(ListNode)
		curNode = curNode.Next
	}
	l2 := new(ListNode)
	curNode = l2
	for i := range 3 {
		curNode.Val = i
		if i == 2 {
			break
		}
		curNode.Next = new(ListNode)
		curNode = curNode.Next
	}
	s1 := addTwoNumbers(l1, l2)
	_ = s1

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	prev := ans
	temp := 0
	carry := 0
	digit := 0
	for l1 != nil || l2 != nil {
		temp = carry
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}

		digit = temp % 10
		temp = temp / 10
		carry = temp

		prev.Next = &ListNode{
			digit,
			nil,
		}
		prev = prev.Next
	}

	if carry != 0 {
		prev.Next = &ListNode{
			carry,
			nil,
		}
	}
	return ans.Next
}
