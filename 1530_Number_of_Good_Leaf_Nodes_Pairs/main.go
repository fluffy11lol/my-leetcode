package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	count := 0
	const MAX_DISTANCE = 10

	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return make([]int, MAX_DISTANCE+1)
		}

		if node.Left == nil && node.Right == nil {
			res := make([]int, MAX_DISTANCE+1)
			res[1] = 1
			return res
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		for i := 1; i <= distance; i++ {
			for j := 1; j <= distance-i; j++ {
				count += left[i] * right[j]
			}
		}

		res := make([]int, MAX_DISTANCE+1)
		for i := 1; i < MAX_DISTANCE; i++ {
			res[i+1] = left[i] + right[i]
		}

		return res
	}

	dfs(root)
	return count
}
