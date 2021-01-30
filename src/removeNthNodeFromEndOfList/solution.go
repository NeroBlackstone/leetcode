package removeNthNodeFromEndOfList

import "github.com/NeroBlackstone/leetcode/collection"

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *collection.ListNode, n int) *collection.ListNode {
	// add head node
	dummy := &collection.ListNode{
		Next: head,
	}
	// find list length
	p, length := head, 0
	for p != nil {
		length++
		p = p.Next
	}
	index := length - n
	// remove Nth node
	p = dummy
	for index > 0 {
		index--
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.2 MB, less than 22.29% of Go online submissions for Remove Nth Node From End of List.
// summary: two pointer, one by one
