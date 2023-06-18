package main

import . "github.com/kovercjm/leetcode-go/data_structure"

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	p, q := head.Next.Next, head.Next
	for ; p != q; p, q = p.Next.Next, q.Next {
		if p == nil || p.Next == nil || p.Next.Next == nil {
			// if reaches nil, means no cycle
			return nil
		}
	}

	for p = head; p != q; p, q = p.Next, q.Next {
	}
	return p
}
