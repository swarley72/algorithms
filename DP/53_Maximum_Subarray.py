class Solution:
    def maxSubArray(self, nums: list[int]) -> int:
        max_sum = nums[0]
        current_sum = nums[0]

        for i in range(1, len(nums)):
            n = nums[i]
            current_sum = max(current_sum + n, n)
            max_sum = max(max_sum, current_sum)

        return max_sum