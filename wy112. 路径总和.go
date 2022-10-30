package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var dfs func(root *TreeNode, sum int) bool
	dfs = func(root *TreeNode, sum int) bool {
		if root != nil {
			if root.Val == sum && root.Left == nil && root.Right == nil {
				return true
			}
			return dfs(root.Left, sum-root.Val) || dfs(root.Right, sum-root.Val)
		}
		return false
	}
	return dfs(root, targetSum)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Println(hasPathSum(root, 22))
}
