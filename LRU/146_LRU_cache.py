class Node:
    def __init__(self, val: int = None, key: int = None):
        self.value = val
        self.next = None
        self.prev = None
        self.key = key


class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.head = Node()
        self.tail = Node()
        self.cache = {}
        self.head.next = self.tail
        self.tail.prev = self.head

    def add_to_head(self, node):
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node

    def remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1

        node = self.cache[key]
        self.remove(node)
        self.add_to_head(node)

        return node.value

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.cache[key]
            node.value = value
            self.remove(node)
            self.add_to_head(node)

            return

        if len(self.cache) == self.capacity:
            last_used_node = self.tail.prev
            self.remove(last_used_node)
            del self.cache[last_used_node.key]

        node = Node(value, key)
        self.cache[key] = node
        self.add_to_head(node)
