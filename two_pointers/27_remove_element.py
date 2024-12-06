def removeElement(nums: list[int], val: int) -> int:
    p = 0
    n = len(nums) - 1

    result = 0

    for i in range(n, -1, -1):
        if nums[i] != val and p <= i:
            print(nums[i],nums)
            while nums[p] != val:
                p += 1
                result += 1
            nums[i], nums[p] = nums[p], nums[i]
            p += 1
            result += 1

    print(result)
    
    return result


removeElement([0,1,2,2,3,0,4,2], 2)

