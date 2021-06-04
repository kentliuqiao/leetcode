package removenthfromendlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	cnt := 0
	for temp != nil {
		cnt++
		temp = temp.Next
	}
	if cnt == 0 {
		return nil
	}
	dummy := &ListNode{Next: head}
	curr := dummy
	for i := 0; i < cnt-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

func removeNthFromEndStack(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{Next: head}
	for tmp := dummy; tmp != nil; tmp = tmp.Next {
		nodes = append(nodes, tmp)
	}
	prev := nodes[len(nodes)-n-1]
	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	return dummy.Next
}

func removeNthFromEndDoublePtr(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
