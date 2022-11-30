package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res = append(res, tmp)
	}
	return res
}
