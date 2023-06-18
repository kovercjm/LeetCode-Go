package main

import (
	"fmt"
	. "github.com/kovercjm/leetcode-go/data_structure"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

func main() {
	//common := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: nil,
	//	},
	//}
	headA := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}
	headB := &ListNode{
		Val:  3,
		Next: nil,
	}
	fmt.Println(getIntersectionNode(headA, headB).Val)
}
