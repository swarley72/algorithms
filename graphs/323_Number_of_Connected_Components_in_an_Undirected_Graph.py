from collections import defaultdict, deque


class Solution:
    def countComponents(self, n: int, edges: list[list[int]]) -> int:
        graph = defaultdict(list)
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)

        res = 0
        seen = set()
        for c in range(n):
            if c not in seen:
                seen.add(c)
                res += 1
                queue = deque([c])
                while queue:
                    node = queue.popleft()

                    for ngh in graph[node]:
                        if ngh not in seen:
                            seen.add(ngh)
                            queue.append(ngh)
        return res
