from collections import defaultdict


class Solution:
    def subarraysDivByK(self, nums: list[int], k: int) -> int:
        prefix_count = defaultdict(int)

        res = 0
        acc = 0
        for n in nums:
            acc += n

            div = acc % k
            if div == 0:
                res += 1
            res += prefix_count.get(div, 0)
            prefix_count[div] += 1
        return res
