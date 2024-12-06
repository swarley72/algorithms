from collections import defaultdict
def threeSum(nums: list[int]) -> list[list[int]]:
    nums = sorted(nums)
    result = []
    n = len(nums)

    for i in range(n):
        target = -nums[i]
        left = i + 1
        right = n - 1

        while left < right:
            current_sum = nums[left] + nums[right]
            
            if current_sum < target:
                left += 1
            elif current_sum > target:
                right -= 1
            else:
                result.append([nums[left], nums[right], target])
                left += 1
                right -= 1

    return result

print(threeSum([-1,0,1]))


print(type(threeSum))
