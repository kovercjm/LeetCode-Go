package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	p, q := dummyHead, head

	for q != nil && q.Next != nil && n > 1 {
		n--
		q = q.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return dummyHead.Next
}
