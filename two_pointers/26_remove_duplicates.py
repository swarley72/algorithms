def remove_duplicates(nums: list[int]) -> int:
    k = 0
    for i in range(len(nums)):
        if nums[k] != nums[i]:
            k += 1
            nums[k], nums[i] = nums[i], nums[k]
    return k + 1
