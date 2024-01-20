package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := head
	slow, fast := head, head
	for i := 1; i < left; i++ {
		pre = slow
		slow = slow.Next
	}
	for i := 1; i <= right; i++ {
		fast = fast.Next
	}
	//node := reverse(slow, fast)
	//pre.Next = node
	//slow.Next = fast
	//return head
	res := reverse(slow, fast)
	slow.Next = fast
	if left == 1 {
		return res
	}
	pre.Next = res
	return head
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

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(reverseBetween(head, 1, 2))
}
