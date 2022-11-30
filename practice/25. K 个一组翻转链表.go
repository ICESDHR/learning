package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	node := reverse(slow, fast)
	slow.Next = reverseKGroup(fast, k)
	return node
}

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

func findKthLargest(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		p := partition(nums, start, end)
		if p > len(nums)-k {
			end = p - 1
		} else if p < len(nums)-k {
			start = p + 1
		} else {
			return nums[p]
		}
	}
	return 0
}
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) []int {
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
	return nums
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	index := start + 1
	for i := index; i <= end; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[start], nums[index-1] = nums[index-1], nums[start]
	return index - 1
}
