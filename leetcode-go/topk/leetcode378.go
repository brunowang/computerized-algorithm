package topk

import (
	"container/heap"
	"github.com/brunowang/computerized-algorithm/leetcode-go/heapq"
)

type Item struct {
	value int
	row   int
	col   int
}

func (i Item) Value() int {
	return i.value
}

func kthSmallest(matrix [][]int, k int) int {
	pq := heapq.PriorityQueue[int, Item]{}
	for i := 0; i < len(matrix); i++ {
		pq = append(pq, Item{value: matrix[i][0], row: i, col: 0})
	}
	heap.Init(&pq)

	var res int
	for k > 0 && pq.Len() > 0 {
		cur := heap.Pop(&pq).(Item)
		res = cur.value
		k--
		row, col := cur.row, cur.col+1
		if col < len(matrix[row]) {
			heap.Push(&pq, Item{value: matrix[row][col], row: row, col: col})
		}
	}
	return res
}
