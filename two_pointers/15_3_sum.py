def three_sum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    result = []
    n = len(nums)

    for i in range(n):
        # пропускаем дубликаты
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        target = -nums[i]
        l = i + 1
        r = n - 1

        while l < r:
            current_sum = nums[l] + nums[r]
            
            if current_sum < target:
                l += 1
            elif current_sum > target:
                r -= 1
            else:
                result.append([nums[l], nums[r], nums[i]])
                l += 1
                r -= 1
                # пропускаем дубликаты слева
                while l < r and nums[l] == nums[l - 1]:
                    l += 1
                # пропускаем дубликаты справа
                while l < r and nums[r] == nums[r + 1]:
                    r += 1

    return result

print(three_sum([-1,0,1]))


print(type(three_sum))
