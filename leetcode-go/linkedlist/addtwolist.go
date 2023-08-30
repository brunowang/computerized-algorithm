package linkedlist

func addTwoList(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	la, lb := reverseList(listA), reverseList(listB)
	result := &ListNode{}
	overflow := 0
	res := result
	for la != nil && lb != nil || overflow != 0 {
		if la == nil {
			la = &ListNode{Val: 0}
		}
		if lb == nil {
			lb = &ListNode{Val: 0}
		}
		res.Next = &ListNode{
			Val: (la.Val + lb.Val + overflow) % 10,
		}
		overflow = (la.Val + lb.Val + overflow) / 10
		la = la.Next
		lb = lb.Next
		res = res.Next
	}
	if la == nil {
		res.Next = lb
	} else if lb == nil {
		res.Next = la
	}
	return reverseList(result.Next)
}
