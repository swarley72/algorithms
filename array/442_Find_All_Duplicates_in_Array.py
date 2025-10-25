class Solution:
    def findDuplicates(self, nums: list[int]) -> list[int]:
        res = []

        for n in nums:
            if nums[abs(n) - 1] < 0:
                res.append(abs(n))
            else:
                nums[abs(n) - 1] = -abs(nums[abs(n) - 1])

        return res