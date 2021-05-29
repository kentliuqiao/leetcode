package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	Cnt  int
}

func Constructor() *MyLinkedList {
	return &MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	sentinel := &ListNode{Next: l.head}
	if sentinel.Next == nil {
		return -1
	}
	node := sentinel.Next
	for i := 0; node != nil; i++ {
		if i == index {
			return node.Val
		}
		node = node.Next
	}
	return -1
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val}
	if l.head == nil {
		l.head = node
		return
	}
	node.Next = l.head
	l.head = node
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{Val: val}
	if l.head == nil {
		l.head = node
		return
	}
	tmp := l.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = node
}

func (l *MyLinkedList) AddAtIndex(val int, index int) {
	sentinel := &ListNode{Next: l.head}
	node := &ListNode{Val: val}
	prev, curr := sentinel, sentinel.Next
	i := 0
	for ; i < index-1 && curr != nil; i++ {
		prev = curr
		curr = curr.Next
	}
	if i != index-1 {
		return
	}
	node.Next = curr
	prev.Next = node
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	sentinel := &ListNode{Next: l.head}
	prev, curr := sentinel, sentinel.Next
	i := 0
	for ; i < index-1 && curr != nil; i++ {
		prev = curr
		curr = curr.Next
	}
	if i != index-1 {
		return
	}
	if curr != nil {
		prev.Next = curr.Next
	} else {
		prev.Next = curr
	}

}
