def get_middle_of_linked_list(head):
    slow = head
    fast = head
    while fast and fast.next:
        fast = fast.next.next
        slow = slow.next
    return slow

def reverse_linked_list(head):
    prev = None
    current = head

    while current:
        tmp = current.next
        current.next = prev
        prev = current
        current = tmp
    return prev



def cycle_linked_list(head):
    slow = head
    fast = head.next

    while fast and fast.next:
        if slow == fast:
            return True
        slow = slow.next
        fast = fast.next.next

    return False
        