package main

import "fmt"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func rightSideView(root *TreeNode1) []int {
	res := []int{}
	trave(root, &res)
	return res
}

func trave(root *TreeNode1, res *[]int) {
	queue := []*TreeNode1{}
	if root == nil {
		return
	}
	queue = append(queue, root)
	for len(queue) > 0 {
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
		*res = append(*res, tmp[len(tmp)-1])
	}
}

func main() {
	root := &TreeNode1{1, nil, nil}
	root.Left = &TreeNode1{2, nil, nil}
	root.Right = &TreeNode1{3, nil, nil}
	root.Left.Right = &TreeNode1{5, nil, nil}
	root.Right.Right = &TreeNode1{4, nil, nil}
	res := rightSideView(root)
	fmt.Println(res)
}
