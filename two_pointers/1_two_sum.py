def two_sum(self, nums: list[int], target: int) -> list[int]:
    hash_map = {}
    for i in range(len(nums)):
        needed = target - nums[i]
        if needed in hash_map:
            return [hash_map[needed], i]
        else:
            hash_map[nums[i]] = i
