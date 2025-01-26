from typing import Optional

class ListNode:
    def __init__(self, val=0):
        self.val = val
        self.next = next


def delete_middle(head: Optional[ListNode]) -> Optional[ListNode]:
    dummy = ListNode()
    dummy.next = head

    slow = dummy
    fast = dummy.next

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    slow.next = slow.next.next
    return dummy.next
