# Есть доска размером M*N.
# В каждой клетке записано целое число.
# Надо расположить шахматную ладью так, чтобы сумма клеток,
# на которые она бьет, была максимальной.
# Клетка, на которую будет стоять, тоже учитывается в сумме.
# Верните эту сумму.
# Напоминаем, что ладья ходит на любое количество клеток по горизонтали и вертикали.
# maxDeadSum:
# * [1, 2, 3]
# * [3, 4, 1]
# * [3, 5, 2]
# Ответ: 16 // максимальная сумма достигается, если ладью поставить в клетку с цифрой 5
# maxDeadSum(
# [
# 	[1, 2, 3],
# 	[3, 4, 1],
# 	[3, 5, 2]
# ]) = 16


def max_dead_sum(board: list[int]): ...


def max_dead_sum_success(board: list[int]):
    n = len(board)
    m = len(board[0])
    rows = [0] * n
    cols = [0] * m
    for i in range(n):
        for j in range(m):
            rows[i] += board[i][j]
            cols[j] += board[i][j]

    max_sum = float("-inf")

    for i in range(n):
        for j in range(m):
            current_sum = rows[i] + cols[j] - board[i][j]
            max_sum = max(max_sum, current_sum)

    return max_sum


assert max_dead_sum_success([[1, 2, 3], [3, 4, 1], [3, 5, 2]]) == 16
assert max_dead_sum_success([[1, 2, 3, 4]]) == 10
