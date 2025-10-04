def removeDuplicates(nums: list[int]) -> int:
	l,r = 0,0

	while r < len(nums):
		cnt = 1
		while r + 1 < len(nums) and nums[r] == nums[r + 1]:
			cnt +=1
			r += 1
		for _ in range(min(2, cnt)):
			nums[l] = nums[r]
			l +=1
		r +=1 
	
	return l
