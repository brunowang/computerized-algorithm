package heapq

import "golang.org/x/exp/constraints"

type Valuer[T constraints.Ordered] interface {
	Value() T
}

type PriorityQueue[T constraints.Ordered, V Valuer[T]] []V

func (q *PriorityQueue[T, V]) Len() int { return len(*q) }

func (q *PriorityQueue[T, V]) Less(i, j int) bool {
	return (*q)[i].Value() < (*q)[j].Value()
}

func (q *PriorityQueue[T, V]) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PriorityQueue[T, V]) Push(x any) {
	v, ok := x.(V)
	if !ok {
		panic("cannot push an element that does not implement ordered interface")
	}
	*q = append(*q, v)
}

func (q *PriorityQueue[T, V]) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
