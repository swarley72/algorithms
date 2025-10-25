class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        lo, hi = 0, len(nums) - 1
        first = -1
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] == target:
                first = mid
                hi = mid - 1
            elif nums[mid] > target:
                hi = mid - 1
            else:
                lo = mid + 1
        
        lo, hi = 0, len(nums) - 1
        last = -1
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] == target:
                last = mid
                lo = mid + 1
            elif nums[mid] > target:
                hi = mid - 1
            else:
                lo = mid + 1

        return [first, last]
                
            