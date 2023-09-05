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

func (n *ListNode) Value() int {
	return n.Val
}

func (n *ListNode) Vals() []int {
	vals := make([]int, 0)
	for i := n; i != nil; i = i.Next {
		vals = append(vals, i.Val)
	}
	return vals
}

func (n *ListNode) Copy() *ListNode {
	h := &ListNode{}
	p := h
	for i := n; i != nil; i = i.Next {
		p.Next = &ListNode{Val: i.Val}
		p = p.Next
	}
	return h.Next
}

func (n *ListNode) Link(node *ListNode, idx int) *ListNode {
	p := node.Find(idx)
	if p == nil {
		return n
	}
	var tail *ListNode
	for i := n; i != nil; i = i.Next {
		tail = i
	}
	tail.Next = p
	return n
}

func (n *ListNode) LinkSelf(idx, nextIdx int) *ListNode {
	node, next := n.Find(idx), n.Find(nextIdx)
	if node != nil && next != nil {
		node.Next = next
	}
	return n
}

func (n *ListNode) Find(idx int) *ListNode {
	p := n
	for i := 0; i < idx; i++ {
		if p == nil {
			return nil
		}
		p = p.Next
	}
	return p
}
