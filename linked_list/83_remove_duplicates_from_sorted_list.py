from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def deleteDuplicates_1(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head
        
        prev = head
        current = head.next

        while current:
            if prev.val == current.val:
                prev.next = current.next
            else:
                prev = current
            current = current.next
        return head

    def deleteDuplicates_2(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head
        
        current = head

        while current and current.next:
            if current.val == current.next.val:
                current.next = current.next.next
            else:
                current = current.next
        return head