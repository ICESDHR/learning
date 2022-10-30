package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	tmp1 := l1
	tmp2 := l2
	head := &ListNode{Val: 0}
	res := head
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			head.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			head.Next = tmp2
			tmp2 = tmp2.Next
		}
		head = head.Next
	}

	if tmp1 != nil {
		head.Next = tmp1
	}

	if tmp2 != nil {
		head.Next = tmp2
	}

	return res.Next
}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 4}

	res := merge(l1, l2)
	fmt.Println(res)
}
