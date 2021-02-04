package collection

//ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewLinkedList : construct linkedlist without head node from int slice
func NewLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	p := head
	for _, e := range nums {
		n := &ListNode{
			Val: e,
		}
		p.Next = n
		p = p.Next
	}
	return head.Next
}

// NewCircularLinkedList : construct linkedlist without head node from int slice and position
// position is index where tail node link next
func NewCircularLinkedList(nums []int,pos int) *ListNode{
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1{
		return NewLinkedList(nums)
	}
	head := &ListNode{}
	p := head
	var posPointer *ListNode
	for i, e := range nums {
		n := &ListNode{
			Val: e,
		}
		if i == pos {
			posPointer=n
		}
		p.Next = n
		p = p.Next
	}
	p.Next=posPointer
	return head.Next
}