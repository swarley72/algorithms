def minimumDifference(nums: list[int], k: int) -> int:
	nums.sort()
	start = 0
	res = float("inf")

	for end in range(len(nums)):
		window_size = end - start + 1
		if window_size == k:
			res = min(res, nums[end] - nums[start])
			start += 1
	
	return res