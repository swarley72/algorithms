def minSubArrayLen(target: int, nums: list[int]) -> int:
    begin = 0
    window_state = 0
    result = float("inf")

    for end in range(len(nums)):
        window_state += nums[end]

        while window_state >= target:
            window_size = end - begin + 1
            result = min(result, window_size)
            window_state -= nums[begin]
            begin += 1

    if result == float("inf"):
        return 0

    return result

print(minSubArrayLen(3, [1, 1, 2, 1, 3, 1, 1]))