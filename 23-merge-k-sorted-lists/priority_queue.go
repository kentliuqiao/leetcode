package mergeksortedlists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := new(PriorityQueue)
	for _, l := range lists {
		if l != nil {
			heap.Push(pq, l)
		}
	}
	// heap.Init(&pq)
	dummy := new(ListNode)
	p := dummy
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}
