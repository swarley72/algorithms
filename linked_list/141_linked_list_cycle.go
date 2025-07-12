type ListNode struct {
	Val  int
	Next *ListNode
}

// вариант с O(n) по времени и O(n) по памяти
func hasCycle(head *ListNode) bool {
	hashMap := make(map[*ListNode]bool)

	current := head

	for current != nil {
		if _, ok := hashMap[current]; ok {
			return true
		}
		hashMap[current] = true
		current = current.Next
	}
	return false
}

// вариант с O(n) по времени и O(1) по памяти, алгоритм флойда
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}