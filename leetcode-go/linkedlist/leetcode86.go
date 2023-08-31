package linkedlist

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。
*/
func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	p, l, r := head, left, right
	for p != nil {
		if p.Val < x {
			l.Next = p
			l = l.Next
		} else {
			r.Next = p
			r = r.Next
		}
		next := p.Next
		p.Next = nil
		p = next
	}
	l.Next = right.Next
	return left.Next
}
