package mergeksortedlists

func mergeTwo(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)

	return mergeTwo(left, right)
}

// 分而治之思想：divideConquer
// 两个有序链表合并为一个有序链表很容易实现，那么可以把原问题——合并K个有序链表分解为多次两个有序链表合并，
// 递归进行下去，直到最后只剩下一个链表即为最终答案。
func divideConquer(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}
