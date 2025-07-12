type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next
		first.Next = second.Next
		second.Next = first
		current.Next = second
		current = first
	}

	return dummy.Next
}