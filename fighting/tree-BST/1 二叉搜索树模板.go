package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 搜索模板
func isInBST(root *TreeNode, num int) bool {
	if root == nil {
		return false
	}
	if root.Val == num {
		return true
	}
	if num < root.Val {
		return isInBST(root.Left, num)
	}
	if num > root.Val {
		return isInBST(root.Right, num)
	}
	return false
}
