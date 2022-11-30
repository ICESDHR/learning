package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	a, b := headA, headB
	for a != nil {
		lenA++
		a = a.Next
	}
	for b != nil {
		lenB++
		b = b.Next
	}

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	}
	if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
