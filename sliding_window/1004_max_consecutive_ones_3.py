def longestOnes(nums: list[int], k: int) -> int:
	begin = 0
	window_state = 0
	result = 0

	for end, n in enumerate(nums):
		if n == 0:
			window_state += 1
		while window_state > k:
			if nums[begin] == 0:
				window_state -= 1
			begin += 1

		result = max(result, end - begin + 1)
	return result
