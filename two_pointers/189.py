# O(n) time, O(n) space
def rotate(nums: list[int], k: int) -> None:
	n = len(nums)
	arr = [0] * n

	for i in range(n):
		new_pos = (i + k) % n
		arr[new_pos] = nums[i]
	
	for i in range(n):
		nums[i] = arr[i]


# O(n) time, O(1) space
def rotate(nums: list[int], k: int) -> None:
	nums.reverse()

	k = k % len(nums)
	
	l, r = 0, k - 1
	while l <r:
		nums[l], nums[r] = nums[r], nums[l]
		l += 1
		r -= 1

	l, r = k, len(nums) - 1
	while l <r:
		nums[l], nums[r] = nums[r], nums[l]
		l += 1
		r -= 1            
        