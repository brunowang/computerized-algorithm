package linkedlist

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
