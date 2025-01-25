def find_max_average( nums: list[int], k: int) -> float:
    begin = 0
    window_state = 0
    result = float("-inf")

    for end in range(len(nums)):
        window_size = end - begin + 1
        window_state += nums[end]
        if window_size == k:
            result = max(result, window_state)
            window_state -= nums[begin]
            begin += 1

    return result / k
