from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def mergeTwoLists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        current = dummy

        first = list1
        second = list2

        while first and second:
            if first.val < second.val:
                current.next = first
                first = first.next
            else:
                current.next = second
                second = second.next
            current = current.next
        
        while first:
            current.next = first
            first = first.next
            current = current.next
        while second:
            current.next = second
            second = second.next
            current = current.next

		# можно проще т.к. если что-то осталось то только в одном списке и поэтому гарантировано в правильном порядке
        if first:
            current.next = first
        if second:
            current.next = second

        return dummy.next