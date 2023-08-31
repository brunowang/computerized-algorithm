package linkedlist

// 删除倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	prev := findNthFromEnd(head, n+1)
	if prev == nil {
		return head.Next
	}
	prev.Next = prev.Next.Next
	return head.Next
}

// 查找倒数第N个节点
func findNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
