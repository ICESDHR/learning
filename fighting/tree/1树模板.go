package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func traverse(root *TreeNode) {
	// 前序
	traverse(root.left)
	// 中序
	traverse(root.right)
	// 后序
}
