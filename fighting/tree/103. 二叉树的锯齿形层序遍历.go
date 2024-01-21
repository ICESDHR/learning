package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	flag := 0
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
		if flag == 1 {
			tmp = swap(tmp)
			flag = 0
		} else {
			flag = 1
		}
		result = append(result, tmp)
	}
	return result
}

func swap(nums []int) []int {
	i, j := 0, len(nums)-1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
