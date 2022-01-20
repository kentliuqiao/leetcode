package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func doublePointer(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode) // 虚拟头节点！
	dummy.Next = head

	p := findKthToLastNode(dummy, n+1) // 找到倒数第N+1个节点，以便删除倒数第N个节点
	if p == nil {
		return head
	}
	p.Next = p.Next.Next
	return dummy.Next
}

func findKthToLastNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for i := 0; i < k; i++ {
		if p == nil {
			return nil
		}
		p = p.Next
	}
	p2 := head
	for p != nil {
		p = p.Next
		p2 = p2.Next
	}
	return p2
}
