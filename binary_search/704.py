
def search(self, nums: list[int], target: int) -> int:
    l,r = 0, len(nums) - 1

    while l <= r:
        middle = l + (r - l) // 2
        if nums[middle] == target:
            return middle
        if nums[middle] > target:
            r = middle - 1
        else:
            l = middle + 1
    
    return -1
        