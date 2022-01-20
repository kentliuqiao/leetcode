package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：走两遍，第一遍计算节点总数N，倒数第K个节点即为整数第N-K个节点。
// 方法二：使用双指针（快慢指针），第一个指针先走K步，然后第二个节点开始同步从头开始走，
// 当第一个节点到达终点时，第二个节点就走了N—K步，也即处于倒数第K个节点。
func findKthToLastNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
