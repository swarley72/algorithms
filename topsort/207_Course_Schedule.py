from collections import defaultdict, deque


class Solution:
    def canFinish(self, numCourses: int, prerequisites: list[list[int]]) -> bool:
        graph = defaultdict(list)
        indegree = defaultdict(int)
        for a, b in prerequisites:
            graph[b].append(a)
            indegree[a] += 1

        queue = deque()
        res = 0
        for i in range(numCourses):
            if indegree[i] == 0:
                queue.append(i)
                res += 1

        while queue:
            course = queue.popleft()

            for ngh in graph[course]:
                indegree[ngh] -= 1
                if indegree[ngh] == 0:
                    res += 1
                    queue.append(ngh)
        return res == numCourses
