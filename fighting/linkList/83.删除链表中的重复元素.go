package main

import "fmt"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func deleteDuplicates(head *ListNode1) *ListNode1 {
	dummy := &ListNode1{}
	result := dummy
	p := head
	mem := make(map[int]struct{})
	for p != nil {
		if _, ok := mem[p.Val]; ok {
			p = p.Next
			continue
		}
		mem[p.Val] = struct{}{}
		dummy.Next = p
		dummy = dummy.Next
	}
	return result.Next
}

func main() {
	head := &ListNode1{1, nil}
	head.Next = &ListNode1{1, nil}
	head.Next.Next = &ListNode1{2, nil}
	head.Next.Next.Next = &ListNode1{3, nil}
	head.Next.Next.Next.Next = &ListNode1{3, nil}
	res := deleteDuplicates(head)
	fmt.Println(res)

}
