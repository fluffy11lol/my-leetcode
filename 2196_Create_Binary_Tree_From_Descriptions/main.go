package main

import "fmt"

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	childrenBranches := make(map[int][]int, 0)
	for _, n := range descriptions {
		arr := childrenBranches[n[0]]
		if len(arr) == 0 {
			arr = make([]int, 2)
		}
		if n[2] == 1 {
			arr[0] = n[1]
		} else {
			arr[1] = n[1]
		}
		childrenBranches[n[0]] = arr
	}

	rootVal := findRoot(descriptions)
	return constructFromRoot(rootVal, childrenBranches)
}

func constructFromRoot(rootVal int, cs map[int][]int) *TreeNode {
	if rootVal == 0 {
		return nil
	}

	children, ok := cs[rootVal]
	if !ok {
		return &TreeNode{
			Val: rootVal,
		}
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  constructFromRoot(children[0], cs),
		Right: constructFromRoot(children[1], cs),
	}
}

func findRoot(descriptions [][]int) int {
	children := make(map[int]bool, 0)

	for _, n := range descriptions {
		children[n[1]] = true
	}

	root := 0

	for _, n := range descriptions {
		found := children[n[0]]
		if !found {
			root = n[0]
			break
		}
	}

	return root
}
