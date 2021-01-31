package swapNodesInPairs

import "github.com/NeroBlackstone/leetcode/collection"

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *collection.ListNode) *collection.ListNode {
	dummyHead := &collection.ListNode{Next: head}
	tempNode := dummyHead
	for tempNode.Next != nil && tempNode.Next.Next != nil {
		preNode, nextNode := tempNode.Next, tempNode.Next.Next
		tempNode.Next = nextNode
		preNode.Next = nextNode.Next
		nextNode.Next = preNode
		tempNode = preNode
	}
	return dummyHead.Next
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
//Memory Usage: 2.1 MB, less than 99.61% of Go online submissions for Swap Nodes in Pairs.