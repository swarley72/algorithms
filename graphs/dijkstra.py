import heapq

# V - количество вершин, E - количество ребер, logV - цена операции с хипой
# сложность по памяти O(V),
# сложность по времени O(V + E) * logV -
def dijkstra(graph, start):

    distances = {}
    min_heap = [(0, graph[start])]

    distances[graph.start] = 0

    while min_heap:
        current_dist, node = heapq.heappop(min_heap)

        if current_dist > distances.get(node, float("inf")):
            continue

        for ngh, dist_to_ngh in graph[node]:
            new_dist = current_dist + dist_to_ngh
            if new_dist < distances.get(ngh, float("inf")):
                distances[ngh] = new_dist
                heapq.heappush(min_heap, (new_dist, ngh))


