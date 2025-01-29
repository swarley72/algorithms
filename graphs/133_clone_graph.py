from typing import Optional
from collections import deque

class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


def clone_graph(self, node: Optional['Node']) -> Optional['Node']:
    if not node:
        return node

    graph = {
        node: Node(node.val, [])
    }

    queue = deque([node])
    while queue:
        v = queue.popleft()

        for ngh in v.neighbors:
            if ngh not in graph:
                graph[ngh] = Node(ngh.val, [])
                queue.append(ngh)
            graph[v].neighbors.append(graph[ngh])

    return graph[node]