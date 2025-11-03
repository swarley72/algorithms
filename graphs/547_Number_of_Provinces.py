from collections import defaultdict, deque


class Solution:
    def findCircleNum(self, isConnected: list[list[int]]) -> int:
        graph = defaultdict(list)
        n = len(isConnected)
        for i in range(n):
            for j in range(n):
                if isConnected[i][j] == 1:
                    graph[i].append(j)
                    graph[j].append(i)

        res = 0
        seen = set()

        for c in range(n):
            if c not in seen:
                res += 1
                seen.add(c)
                queue = deque([c])
                while queue:
                    node = queue.popleft()

                    for ngh in graph[node]:
                        if ngh not in seen:
                            seen.add(ngh)
                            queue.append(ngh)
        return res
