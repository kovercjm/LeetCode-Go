package main

import . "github.com/kovercjm/leetcode-go/data_structure"

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	for p, q := head, head.Next; q != nil; p, q = p.Next, q.Next.Next {
		if p == q {
			return true
		}
		if q.Next == nil {
			return false
		}
	}
	return false
}
