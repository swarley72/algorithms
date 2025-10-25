class Solution:
    def singleNumber(self, nums: list[int]) -> int:
        mask = 0

        for n in nums:
            mask ^= n

        return mask