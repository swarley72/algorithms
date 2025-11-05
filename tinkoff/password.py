# Перед входом в офис стоит кодовый замок состоящий из кнопок с числами,
# например:
# 1   2   3   4
# 5   6   7   8
# 9  10  11  12
# Кодовый замок может быть любого размера,
# но он всегда прямоугольной или квадратной формы

# Каждый день пароль меняется и пользователям присылается пароль на текущий день в виде массива чисел.
# Необходимо определить, сколько движений пальца потребуется,
# чтобы набрать весь пароль. Двигать пальцем можно вверх, вниз, вправо, влево и по диагонали.

# Пример: для пароля [1, 9, 3, 7] потребуется 5 движений пальца (1 → 9 - два движения вниз, 9 → 3 два движения по диагонали, 3 → 7 одно движение вниз).

from collections import deque


def get_moves(grid: list[list[int]], password: list[int]): ...


assert (
    get_moves(
        [
            [1, 2, 3, 4],
            [5, 6, 7, 8],
            [9, 10, 11, 12],
        ],
        [1, 9, 3, 7],
    )
) == 5

assert get_moves([[1, 2, 3], [4, 5, 6], [7, 8, 9]], [1, 5, 9]) == 2
assert get_moves([[1, 2, 3]], [2]) == 0
assert (
    get_moves(
        [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]], [1, 16, 13, 4]
    )
    == 9
)
assert get_moves([[1, 2, 3, 4, 5]], [1, 5, 3]) == 6


def get_moves_success(grid: list[list[int]], password: list[int]):
    if not password:
        return 0
    n = len(grid)
    m = len(grid[0])
    res = 0

    start = None
    for i in range(n):
        for j in range(m):
            if grid[i][j] == password[0]:
                start = (i, j)
                break

    def is_valid(r, c, visited) -> bool:
        return r >= 0 and r < n and c >= 0 and c < m and (r, c) not in visited

    for k in password:
        queue = deque([(start, 0)])
        visited = set(start)
        while queue:
            coord, steps = queue.popleft()
            r, c = coord
            if grid[r][c] == k:
                res += steps
                start = (r, c)
                break

            for dx, dy in [
                (1, 0),
                (0, 1),
                (-1, 0),
                (0, -1),
                (1, 1),
                (1, -1),
                (-1, 1),
                (-1, -1),
            ]:
                row = dx + r
                col = dy + c
                if is_valid(row, col, visited):
                    queue.append(((row, col), steps + 1))
                    visited.add((row, col))
    return res
