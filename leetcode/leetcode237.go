package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// what should I do when node is the last node in ListNode
	if node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	p := node.Next
	node.Next = node.Next.Next
	p.Next = nil
}
