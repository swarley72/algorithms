# Даны три неубывающих массива чисел.
# Найти число, которое присутствует во всех трех массивах.

# Input: [1,2,4,5], [3,3,4], [2,3,4,5,6]
# Output: 4


def get_number(arr1: list[int], arr2: list[int], arr3: list[int]) -> int: ...


assert get_number([1, 2, 4, 5], [3, 3, 4], [2, 3, 4, 5, 6]) == 4


# Число в начале
assert get_number([1, 2, 3], [1, 4, 5], [1, 6, 7]) == 1
# Число в конце
assert get_number([1, 2, 5], [3, 4, 5], [4, 5, 6]) == 5

# Нет общего числа
assert get_number([1, 2, 3], [4, 5, 6], [7, 8, 9]) == -1

# Все массивы одинаковые
assert get_number([1, 2, 3], [1, 2, 3], [1, 2, 3]) == 1

# Дубликаты
assert get_number([1, 1, 2, 3], [1, 2, 2, 3], [1, 2, 3, 3]) == 1
# Проблемный кейс из анализа
assert get_number([1, 2, 3], [2, 2, 3], [1, 3, 4]) == 3
# Один элемент в каждом
assert get_number([5], [5], [5]) == 5
assert get_number([1], [2], [3]) == -1
# Разные длины
assert get_number([1, 2, 3, 4, 5], [3], [1, 2, 3, 4]) == 3


def get_number_success(arr1: list[int], arr2: list[int], arr3: list[int]) -> int:
    p1 = 0
    p2 = 0
    p3 = 0

    while p1 < len(arr1) and p2 < len(arr2) and p3 < len(arr3):
        num_1 = arr1[p1]
        num_2 = arr2[p2]
        num_3 = arr3[p3]
        if num_1 == num_2 and num_1 == num_3:
            return num_1

        min_val = min(num_1, num_2, num_3)

        if num_1 == min_val:
            p1 += 1
        if num_2 == min_val:
            p2 += 1
        if num_3 == min_val:
            p3 += 1

    return -1
