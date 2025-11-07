from collections import deque


class Solution:
    def shortestPathBinaryMatrix(self, grid: list[list[int]]) -> int:
        if grid[0][0] == 1:
            return -1
        m = len(grid)
        n = len(grid[0])
        queue = deque([(0, 0, 1)])
        grid[0][0] = 1

        while queue:
            r, c, steps = queue.popleft()
            if r == m - 1 and c == n - 1:
                return steps

            for dx, dy in [
                (1, 0),
                (0, 1),
                (-1, 0),
                (0, -1),
                (1, 1),
                (-1, 1),
                (1, -1),
                (-1, -1),
            ]:
                row = r + dx
                col = c + dy
                if (
                    row >= 0
                    and row < m
                    and col >= 0
                    and col < n
                    and grid[row][col] == 0
                ):
                    grid[row][col] = 1
                    queue.append((row, col, steps + 1))
        return -1
