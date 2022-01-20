package intersectionoftwolists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路一：两个指针一起从两个链表头节点走，每个指针走到链表结尾时，换到另外一个链表上，直到二者相遇或者同时到达链表结尾（表示未相遇）
// 图示：getIntersectionNode
// a1 -> a2 -> c -> a3 -> a4 -> b1 -> b2 -> b3 -> c
// b1 -> b2 -> b3 -> c -> b4 -> b5 -> a1 -> a2 -> c
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	l1, l2 := 0, 0
	for a != nil {
		a = a.Next
		l1++
	}
	for b != nil {
		b = b.Next
		l2++
	}
	a, b = headA, headB
	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			a = a.Next
		}
	} else {
		for i := 0; i < l2-l1; i++ {
			b = b.Next
		}
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}
