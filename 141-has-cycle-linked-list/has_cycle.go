package hascyclelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodes[head]; ok {
			return true
		}
		nodes[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycleFastSlowPointer(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		// 快指针到达了链表尾部
		if fast == nil || fast.Next == nil {
			return false
		}
		// 慢指针每次走一步
		slow = slow.Next
		// 快指针每次走两步
		fast = fast.Next.Next
	}
	// 快指针套圈了慢指针，则表示有环
	return true
}
