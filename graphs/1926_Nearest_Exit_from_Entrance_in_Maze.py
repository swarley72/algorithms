from collections import deque


class Solution:
    def nearestExit(self, maze: list[list[str]], entrance: list[int]) -> int:
        m = len(maze)
        n = len(maze[0])
        maze[entrance[0]][entrance[1]] = "+"

        queue = deque([(entrance[0], entrance[1], 0)])
        while queue:
            r, c, steps = queue.popleft()

            if (r == 0 or c == 0 or r == m - 1 or c == n - 1) and (
                r != entrance[0] or c != entrance[1]
            ):
                return steps

            for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                row = r + dx
                col = c + dy
                if (
                    row >= 0
                    and row < m
                    and col >= 0
                    and col < n
                    and maze[row][col] == "."
                ):
                    maze[row][col] = "+"
                    queue.append((row, col, steps + 1))

        return -1
