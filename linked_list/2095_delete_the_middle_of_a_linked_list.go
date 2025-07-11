type ListNode struct {
	Val  int
	Next *ListNode
}

// по условию head всегда есть
func deleteMiddle(head *ListNode) *ListNode {
	// проверяем что есть 2 ноды
	if head.Next == nil {
		// деволтное значение указателя nil
		return nil
	}
	slow := head
	// даем фору
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
