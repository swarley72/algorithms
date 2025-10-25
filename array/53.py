def maxSubArray(self, nums: list[int]) -> int:
    res = nums[0]

    current_max = nums[0]

    for i in range(1, len(nums)):
        n = nums[i]

        current_max = max(current_max + n, n)
        res = max(res, current_max)

    return res
