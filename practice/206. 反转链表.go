package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归方法
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

// 非递归方法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head
	for cur != nil {
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
	head.Next.Next = &ListNode{Val: 3}
	fmt.Println(reverseList(head))
}

// 反转a到b之间的链表，包含a，不包含b
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
