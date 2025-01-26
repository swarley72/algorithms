import heapq
import math

# не меняем знак при вытаскивании из кучи
def min_stone_sum_floor(piles: list[int], k: int) -> int:
    max_heap = [-s for s in piles]
    heapq.heapify(max_heap)

    for _ in range(k):
        stone = heapq.heappop(max_heap)
        new_stones = math.floor(stone / 2)
        heapq.heappush(max_heap, new_stones)

    return -sum(max_heap)

# меняем знак при вытаскивании из кучи
def min_stone_sum_ceil(piles: list[int], k: int) -> int:
    max_heap = [-s for s in piles]
    heapq.heapify(max_heap)

    for _ in range(k):
        stone = -heapq.heappop(max_heap)
        new_stones = math.ceil(stone / 2)
        heapq.heappush(max_heap, -new_stones)

    return -sum(max_heap)