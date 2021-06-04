package swappairslinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		left := temp.Next
		right := temp.Next.Next
		temp.Next = right
		left.Next = right.Next
		right.Next = left
		temp = left
	}
	return dummy.Next
}

func swap(first, second *ListNode) *ListNode {
	first.Next = second.Next
	second.Next = first
	return second
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	swapped := swapPairsRecursive(newHead.Next)
	head.Next = swapped
	newHead.Next = head
	return newHead
}
