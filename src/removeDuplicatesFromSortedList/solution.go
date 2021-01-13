package removeDuplicatesFromSortedList

import "github.com/NeroBlackstone/leetcode/collection"

//https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *collection.ListNode) *collection.ListNode {
	now := head
	// list that have at least 2 nodes will execute this for loop
	for now != nil && now.Next != nil {
		if now.Next.Val == now.Val {
			now.Next = now.Next.Next
		} else {
			now = now.Next
		}
	}
	return head
}

// Runtime: 4 ms, faster than 85.61% of Go online submissions for Remove Duplicates from Sorted List.
//Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
