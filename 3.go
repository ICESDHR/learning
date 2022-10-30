package main
package main

import (
"fmt"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool{
	if root == nil{
		return false
	}
	var dfs func(root *TreeNode, sum int) bool
	dfs = func(root *TreeNode, sum int) bool{
		if root != nil{
			if root.Val == sum && root.Left==nil && root.Right==nil{
				return true
			}
			dfs(root.Left, sum-root.Val)
			dfs(root.Right, sum-root.Val)
		}
		return false
	}
	return dfs(root, sum)
}

