package main

import "math"

type TreeNode124 struct {
	Val   int
	Left  *TreeNode124
	Right *TreeNode124
}

var result int

func maxPathSum(root *TreeNode124) int {
	result = math.MinInt
	traverse(root)
	return result
}

func traverse(root *TreeNode124) int {
	if root == nil {
		return 0
	}
	left := max(0, traverse(root.Left))
	right := max(0, traverse(root.Right))
	result = max(result, root.Val+left+right)

	return max(left, right) + root.Val
}

func max(nums ...int) int {
	maxNum := math.MinInt
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
