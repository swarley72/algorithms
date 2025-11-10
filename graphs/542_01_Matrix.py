from collections import deque


class Solution:
    def updateMatrix(self, mat: list[list[int]]) -> list[list[int]]:
        m = len(mat)
        n = len(mat[0])

        queue = deque()
        for i in range(m):
            for j in range(n):
                if mat[i][j] == 0:
                    queue.append((i, j, 0))
        seen = set()
        while queue:
            r, c, steps = queue.popleft()

            for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                row = r + dx
                col = c + dy
                if (
                    row >= 0
                    and row < m
                    and col >= 0
                    and col < n
                    and (row, col) not in seen
                    and mat[row][col] != 0
                ):
                    seen.add((row, col))
                    queue.append((row, col, steps + 1))
                    mat[row][col] = steps + 1

        return mat
