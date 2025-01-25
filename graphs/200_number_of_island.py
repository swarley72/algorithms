def num_of_islands_no_space(grid: list[list[str]]) -> int:
    m = len(grid)
    n = len(grid[0])

    result = 0

    def is_valid(r: int, c: int):
        return 0 <= r < m and 0 <= c < n and grid[r][c] == "1"

    for i in range(m):
        for j in range(n):
            if grid[i][j] == "1":
                result += 1
                grid[i][j] = "0"
                stack = [(i, j)]
                while stack:
                    r, c = stack.pop()

                    for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                        row = r + dx
                        col = c + dy
                        if is_valid(row, col):
                            grid[row][col] = "0"
                            stack.append((row, col))

    return result

def num_of_islands_with_extra_space(grid: list[list[str]]) -> int:
    m = len(grid)
    n = len(grid[0])

    result = 0
    visited = set()

    def is_valid(r, c):
        return 0 <= r < m and 0 <= c < n and grid[r][c] == "1" and (r, c) not in visited

    for i in range(m):
        for j in range(n):
            if grid[i][j] == "1" and (i, j) not in visited:
                result += 1
                visited.add((i, j))
                stack = [(i, j)]
                while stack:
                    r, c = stack.pop()

                    for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                        row = r + dx
                        col = c + dy
                        if is_valid(row, col):
                            visited.add((row, col))
                            stack.append((row, col))

    return result
