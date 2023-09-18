package linkedlist

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	vhead := &ListNode{Next: head}
	prev := vhead
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	tail := prev
	for i := 0; i < right-left+1; i++ {
		tail = tail.Next
	}
	remain := tail.Next
	tail.Next = nil
	list := reverseList(prev.Next)
	prev.Next.Next = remain
	prev.Next = list
	return vhead.Next
}
