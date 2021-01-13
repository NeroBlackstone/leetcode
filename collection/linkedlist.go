package collection

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// construct linklist without head node from int slice
func NewLinkList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	l := &ListNode{
		Val: s[0],
	}
	p := l
	for _, e := range s[1:] {
		n := &ListNode{
			Val: e,
		}
		p.Next = n
		p = p.Next
	}
	return l
}
