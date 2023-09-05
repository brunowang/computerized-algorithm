package linkedlist

import (
	"container/heap"
	"github.com/brunowang/computerized-algorithm/leetcode-go/heapq"
)

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	vhead := &ListNode{}
	p := vhead
	pq := make(heapq.PriorityQueue[int, *ListNode], len(lists))
	for i, head := range lists {
		if head != nil {
			pq[i] = head
		}
	}
	heap.Init(&pq)

	for pq.Len() != 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		p = p.Next
	}
	return vhead.Next
}
