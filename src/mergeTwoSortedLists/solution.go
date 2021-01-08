package mergeTwoSortedLists

import (
	"github.com/NeroBlackstone/leetcode/collection"
)

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *collection.ListNode, l2 *collection.ListNode) *collection.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	p1, p2 := l1, l2
	// init Head Node
	pResult := &collection.ListNode{}
	pr := pResult
	for p1 != nil && p2 != nil {
		p1Val, p2Val := p1.Val, p2.Val
		if p1Val <= p2Val {
			pr.Next = p1
			p1 = p1.Next
		} else {
			pr.Next = p2
			p2 = p2.Next
		}
		pr = pr.Next
	}
	if p1 == nil {
		pr.Next = p2
	}
	if p2 == nil {
		pr.Next = p1
	}
	return pResult.Next
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.5 MB, less than 46.77% of Go online submissions for Merge Two Sorted Lists.
// 总结：题目里的constraints不需要在解题中验证，一般不会给边界条件的测试数据，不要在这上面浪费时间，如果碰上特殊输入再验证。
