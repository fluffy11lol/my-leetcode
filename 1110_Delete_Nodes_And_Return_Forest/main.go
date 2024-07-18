package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDeleteMap := make(map[int]bool)
	for _, val := range to_delete {
		toDeleteMap[val] = true
	}

	var forest []*TreeNode

	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		shouldDelete := toDeleteMap[node.Val]

		if isRoot && !shouldDelete {
			forest = append(forest, node)
		}

		node.Left = dfs(node.Left, shouldDelete)
		node.Right = dfs(node.Right, shouldDelete)

		if shouldDelete {
			return nil
		}
		return node
	}

	dfs(root, true)
	return forest
}
