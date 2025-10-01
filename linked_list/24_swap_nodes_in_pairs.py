from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def swapPairs(head: Optional[ListNode]) -> Optional[ListNode]:
	if not head or not head.next:
		return head
	
	
	dummy = ListNode()
	dummy.next = head

	current = dummy

	while current.next and current.next.next:
		first = current.next
		second = current.next.next
		first.next = second.next
		second.next = first
		current.next = second
		current = first
	
	return dummy.next
