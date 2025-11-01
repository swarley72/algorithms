class Solution:
    def summaryRanges(self, nums: list[int]) -> list[str]:
        if not nums:
            return []

        if len(nums) == 1:
            return [str(nums[0])]

        res = []
        first = nums[0]
        prev = nums[0]
        for i in range(1, len(nums)):
            if prev + 1 == nums[i]:
                prev = nums[i]
            elif first == prev:
                res.append(f"{first}")
                first = nums[i]
                prev = nums[i]
            else:
                res.append(f"{first}->{prev}")
                first = nums[i]
                prev = nums[i]

        if first == prev:
            res.append(f"{first}")
        else:
            res.append(f"{first}->{prev}")

        return res
