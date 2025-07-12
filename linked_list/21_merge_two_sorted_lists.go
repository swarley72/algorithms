type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	first := list1
	second := list2
	dummy := &ListNode{}
	current := dummy

	for first != nil && second != nil {
		if first.Val > second.Val {
			current.Next = second
			second = second.Next
		} else {
			current.Next = first
			first = first.Next
		}

		current = current.Next
	}

	// можно проще
	for first != nil {
		current.Next = first
		first = first.Next
		current = current.Next
	}
	for second != nil {
		current.Next = second
		second = second.Next
		current = current.Next
	}

	// потому что если что-то осталось то только в одном списке
	// а так как он уже отсортирован мы просто указываем на то место где остановились
	// и там гарантировано значения больше и в правильном порядке
	if first != nil {
		current.Next = first
	}

	if first != nil {
		current.Next = second
	}

	return dummy.Next
}