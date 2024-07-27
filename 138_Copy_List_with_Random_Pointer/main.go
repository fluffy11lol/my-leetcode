package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr, hashMap := head, make(map[*Node]*Node)
	hashMap[nil] = nil

	for curr != nil {
		newNode := new(Node)
		newNode.Val = curr.Val
		hashMap[curr] = newNode
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		newNode := hashMap[curr]
		newNode.Next = hashMap[curr.Next]
		newNode.Random = hashMap[curr.Random]
		curr = curr.Next
	}

	return hashMap[head]
}
