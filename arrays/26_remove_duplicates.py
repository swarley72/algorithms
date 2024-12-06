def remove_duplicates(nums: list[int]) -> int:
    curr = nums[0]
    count = 0
    for i in range(1, len(nums)):
        if nums[i] == curr:
            count += 1
            nums[i] = "_"
        elif count > 0:
            nums[i - count] = nums[i]
            nums[i] = "_"
            curr = nums[i]
        else:
            curr = nums[i]
        
        print(curr, count, nums)

remove_duplicates([0, 0, 1, 1, 1, 2, 2, 2, 3])

print(10 ** 3)
