package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	newHead.Next, head.Next = head, swapPairs(head.Next.Next)
	return newHead
}
