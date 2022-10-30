package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		length := len(queue)
		level := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		if flag {
			reverse(level)
			res = append(res, level)
			flag = false
		} else {
			res = append(res, level)
			flag = true
		}

	}
	return res
}

func reverse(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(zigzagLevelOrder(root))
}
