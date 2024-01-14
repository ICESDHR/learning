package main

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

func levelOrder(root *TreeNode102) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode102{}
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
		result = append(result, tmp)
	}
	return result
}
