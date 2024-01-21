package linkList

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:m/2])
	right := mergeKLists(lists[m/2:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(left, right *ListNode) *ListNode {
	pre := &ListNode{}
	tmp := pre
	for left != nil && right != nil {
		if left.Val < right.Val {
			tmp.Next = left
			left = left.Next
		} else {
			tmp.Next = right
			right = right.Next
		}
		tmp = tmp.Next
	}
	if left != nil {
		tmp.Next = left
	}
	if right != nil {
		tmp.Next = right
	}
	return pre.Next
}
