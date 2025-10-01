def moveZeroes(nums: list[int]) -> None:
	n = len(nums)
	p1 = 0
	# доходим до первого нуля
	while p1 <= n - 1 and nums[p1] != 0:
		p1 += 1
	
	for i in range(p1 + 1, n):
		if nums[i] != 0:
			nums[p1], nums[i] = nums[i], nums[p1]
			p1 += 1
	
        