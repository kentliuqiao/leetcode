package detectcyclelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	nodes := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodes[head]; ok {
			return head
		}
		nodes[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycleFastSlowPtr(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
