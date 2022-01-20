package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func doublePointer(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	q1, q2 := list1, list2
	for q1 != nil && q2 != nil {
		if q1.Val < q2.Val {
			p.Next = q1
			q1 = q1.Next
		} else {
			p.Next = q2
			q2 = q2.Next
		}
		p = p.Next
	}
	if q1 != nil {
		p.Next = q1
	}
	if q2 != nil {
		p.Next = q2
	}
	return dummy.Next
}
