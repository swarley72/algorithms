from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode()
        dummy.next = head
        
        slow = dummy
        fast = dummy
        for _ in range(n):
            fast = fast.next

        while fast and fast.next:
            slow = slow.next
            fast = fast.next
        slow.next = slow.next.ne

        return dummy.next