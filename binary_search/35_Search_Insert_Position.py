class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        lo = 0
        hi = len(nums) - 1
        res = -1

        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                res = mid
                hi = mid - 1
            else:
                lo = mid + 1
        
        if res == -1:
            return len(nums)

        return res