package linkedlist

/*
给你单链表的头结点 head ，请你找出并返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。
*/
func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return p2
}
