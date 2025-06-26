def two_sum_sorted(nums: list[int], target: int) -> list[int]:
    l = 0
    r = len(nums) - 1
    while True:
        two_sum = nums[l] + nums[r]
        if two_sum == target:
            return [l + 1, r + 1]

        if two_sum > target:
            r -= 1
        else:
            l += 1
