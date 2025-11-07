import heapq


class Solution:
    def lastStoneWeight(self, stones: list[int]) -> int:
        min_heap = [-s for s in stones]
        heapq.heapify(min_heap)

        while len(min_heap) > 1:
            first = heapq.heappop(min_heap)
            second = heapq.heappop(min_heap)
            if first == second:
                continue
            heapq.heappush(min_heap, -(abs(first) - abs(second)))
        if not min_heap:
            return 0
        return -min_heap[0]
