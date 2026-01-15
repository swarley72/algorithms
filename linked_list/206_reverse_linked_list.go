type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// если сделать тут просто prev := &ListNode{}
	var prev *ListNode
	current := head
	for current != nil {
		tmpNext := current.Next
		current.Next = prev
		prev = current
		current = tmpNext
	}

	return prev
}