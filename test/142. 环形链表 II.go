package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	hasCycle, count := hasCycle(head)
	if !hasCycle {
		return nil
	}
	length := 0
	slow := head
	fast := head.Next
	for i := 0; i < count; i++ {
		fast = fast.Next
	}
	for slow != fast {
		length++
		slow = slow.Next
		fast = fast.Next
	}
	res := head
	for i := 0; i < length; i++ {
		res = res.Next
	}
	return res
}

//func hasCycle(head *ListNode) (bool, int) {
//	slow := head
//	fast := head.Next
//	count := 0
//	for slow != fast && fast.Next != nil {
//		count++
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	if slow == fast {
//		return true, count
//	} else {
//		return false, 0
//	}
//}

func hasCycle(head *ListNode) (bool, int) {
	slow := head
	fast := head
	count := 0
	for fast != nil && fast.Next != nil {
		count++
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == fast {
		return true, count
	} else {
		return false, 0
	}
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = head
	fmt.Println(detectCycle(head))
}
