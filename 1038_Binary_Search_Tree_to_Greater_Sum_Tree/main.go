package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	bstToGstR(root, 0)
	return root
}

func bstToGstR(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	rightSum := bstToGstR(root.Right, sum)
	root.Val += rightSum
	leftSum := bstToGstR(root.Left, root.Val)
	return leftSum
}
