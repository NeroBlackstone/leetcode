package linkedListCycle

import "github.com/NeroBlackstone/leetcode/collection"

//https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *collection.ListNode) bool {
    if head==nil||head.Next==nil{
		return false
	}
	slow,fast:=head,head.Next
	for fast!=slow{
		if fast==nil||fast.Next==nil{
			return false
		}
		slow,fast=slow.Next,fast.Next.Next
	}
	return true
}

//Runtime: 16 ms, faster than 5.58% of Go online submissions for Linked List Cycle.
//Memory Usage: 4.6 MB, less than 23.57% of Go online submissions for Linked List Cycle.