from collections import defaultdict

class Solution:
    def subarraySum(self, nums: list[int], k: int) -> int:
        prefix_sum = 0
        count = 0
        prefix_count = defaultdict(int)
        # prefix_count = {0: 1} # важно! для подмассивов от начала
        prefix_count[0] += 1

        for n in nums:
            prefix_sum += n

            # Ищем prefix_sum - k в хэш-таблице
            if prefix_sum - k in prefix_count: 
                count += prefix_count[prefix_sum - k]

            # Сохраняем текущую prefix_sum
            prefix_count[prefix_sum] += 1

        return count