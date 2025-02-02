from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
    if not head:
        return head

    current = head
    prev = None
    while current:
        old_next = current.next
        current.next = prev
        prev = current
        current = old_next

    return prev
