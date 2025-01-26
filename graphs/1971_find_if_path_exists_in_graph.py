from collections import defaultdict, deque

def valid_path_stack(n: int, edges: list[list[int]], source: int, destination: int) -> bool:
    graph = defaultdict(list)

    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    visited = set()
    visited.add(source)

    stack = [source]
    while stack:
        v = stack.pop()

        if v == destination:
            return True

        for ngh in graph[v]:
            if ngh not in visited:
                visited.add(ngh)
                stack.append(ngh)

    return False

def valid_path_queue(n: int, edges: list[list[int]], source: int, destination: int) -> bool:
    graph = defaultdict(list)

    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    visited = set()
    visited.add(source)

    queue = deque([source])
    while queue:
        v = queue.popleft()

        if v == destination:
            return True

        for ngh in graph[v]:
            if ngh not in visited:
                visited.add(ngh)
                queue.append(ngh)

    return False
