import heapq
from collections import defaultdict

def network_delay_time(times: list[list[int]], n: int, k: int) -> int:
    graph = defaultdict(list)

    for a, b, w in times:
        graph[a].append((b, w))

    distances = {k: 0}

    min_heap = [(0, k)]

    while min_heap:
        cost, node = heapq.heappop(min_heap)

        for ngh, cost_to_ngh in graph[node]:
            new_cost = cost + cost_to_ngh
            if new_cost < distances.get(ngh, float("inf")):
                distances[ngh] = new_cost
                heapq.heappush(min_heap, (new_cost, ngh))

    if len(distances) != n:
        return -1

    return max(distances.values())
