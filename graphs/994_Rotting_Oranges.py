from collections import deque


class Solution:
    def orangesRotting(self, grid: list[list[int]]) -> int:
        m = len(grid)
        n = len(grid[0])

        count_orange = 0
        rotten = 0
        queue = deque()

        for i in range(m):
            for j in range(n):
                if grid[i][j] != 0:
                    count_orange += 1
                if grid[i][j] == 2:
                    rotten += 1
                    queue.append((i, j, 0))
        steps = 0

        while queue:
            r, c, step = queue.popleft()
            steps = max(steps, step)

            for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                row = r + dx
                col = c + dy

                if (
                    row >= 0
                    and row < m
                    and col >= 0
                    and col < n
                    and grid[row][col] == 1
                ):
                    rotten += 1
                    grid[row][col] = 2
                    queue.append((row, col, step + 1))

        if count_orange != rotten:
            return -1

        return steps
