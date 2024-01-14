package main

import "math"

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}
