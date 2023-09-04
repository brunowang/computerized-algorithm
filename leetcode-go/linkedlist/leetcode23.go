package linkedlist

import "container/heap"

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
	pq := make(Queue, len(lists))
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

type Queue []*ListNode

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*ListNode))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
