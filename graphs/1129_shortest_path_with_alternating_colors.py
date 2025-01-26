from collections import defaultdict, deque

def shortest_alternating_paths(self, n: int, redEdges: list[list[int]], blueEdges: list[list[int]]) -> list[int]:
    RED = 1
    BLUE = 0
    graph = defaultdict(lambda: defaultdict(list))
    color_map = {
        RED: BLUE,
        BLUE: RED,
    }

    for a, b in redEdges:
        graph[RED][a].append(b)

    for a, b in blueEdges:
        graph[BLUE][a].append(b)

    result = [float("inf")] * n
    result[0] = 0
    visited = set()
    visited.add((0, BLUE))
    visited.add((0, RED))

    # 0 - нода, 1 - цвет, 2 - шаги
    queue = deque([(0, BLUE, 0), (0, RED, 0)])

    while queue:
        node, color, steps = queue.popleft()
        result[node] = min(result[node], steps)

        for ngh in graph[color][node]:
            next_color = color_map[color]
            if (ngh, next_color) not in visited:
                visited.add((ngh, next_color))
                queue.append((ngh, next_color, steps + 1))

    return [x if x != float("inf") else -1 for x in result]
