from dataclasses import dataclass
from typing import Optional


@dataclass
class Node:
    value: int
    next: Optional["Node"] = None

@dataclass
class MyLinkedListClassic:
    size: int = 0
    head: Node | None = None
    
    def add_head(self, val: int):
        self.add_at_index(0, val)

    def add_tail(self, val: int):
        self.add_at_index(self.size, val)

    def get_by_index(self, index: int):
        if index < 0 or index >= self.size:
            return

        current = self.head        
        for _ in range(index):
            current = current.next
        
        return current.value

    
    def add_at_index(self, index: int, val: int):
        if index < 0 or index > self.size:
            return
        
        self.size += 1
        new_node = Node(val)

        if index == 0:
            new_node.next = self.head
            self.head = new_node
        else:
            current = self.head
            for _ in range(index - 1):
                current = current.next
            old_next = current.next
            current.next = new_node
            new_node.next = old_next
    
    def delete_by_index(self, index: int):
        ...
        