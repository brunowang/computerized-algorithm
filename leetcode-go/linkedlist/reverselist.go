package linkedlist

/*
leetcode206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		curr, prev = swapListNode(curr, prev)
	}
	return prev
}

func swapListNode(curr *ListNode, prev *ListNode) (*ListNode, *ListNode) {
	next := curr.Next
	curr.Next = prev
	return next, curr
}

// backup code
func reverseList0(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	var next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
