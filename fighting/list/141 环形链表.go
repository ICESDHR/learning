package main

func hasCycle(head *ListNode) bool {
	track := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := track[head]; ok {
			return true
		}
		track[head] = struct{}{}
		head = head.Next
	}
	return false
}
