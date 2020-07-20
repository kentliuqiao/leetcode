package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	total := 0
	dummy := &ListNode{}
	dummy.Next = head
	p := head
	for p != nil {
		total++
		p = p.Next
	}
	p = dummy
	for i := 0; i < total-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}

// 双指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p, q := head, dummy
	for i := 0; i < n+1; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}
