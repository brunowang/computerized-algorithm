package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals ...int) *ListNode {
	h := &ListNode{}
	p := h
	for _, v := range vals {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return h.Next
}

func (n *ListNode) Vals() []int {
	vals := make([]int, 0)
	for i := n; i != nil; i = i.Next {
		vals = append(vals, i.Val)
	}
	return vals
}
