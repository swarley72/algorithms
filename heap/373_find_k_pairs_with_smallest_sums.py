import heapq

def k_smallest_pairs(self, nums1: list[int], nums2: list[int], k: int) -> list[list[int]]:
    min_heap = [[nums1[0] + nums2[0], 0, 0]]
    result = []
    visited = set()
    visited.add((0, 0))

    while min_heap and len(result) < k:
        _, i, j = heapq.heappop(min_heap)
        result.append([nums1[i], nums2[j]])

        if i + 1 < len(nums1) and (i + 1, j) not in visited:
            heapq.heappush(min_heap, [nums1[i + 1] + nums2[j], i + 1, j])
            visited.add((i + 1, j))

        if j + 1 < len(nums2) and (i, j + 1) not in visited:
            heapq.heappush(min_heap, [nums1[i] + nums2[j + 1], i, j + 1])
            visited.add((i, j + 1))

    return result

# временная сложность - O(k log k)


# пространственная сложность O(k)
# Храним k результатов: O(k).
# Храним k элементов в куче: O(k).
# Храним visited (уникальные пары): максимум O(k), так как не добавляем больше k пар.
