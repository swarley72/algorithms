# time - O(n) space - O(n)
def findDisappearedNumbers(nums: list[int]) -> list[int]:
	hash_map = set(nums)

	res = []
	for n in range(1, len(nums) + 1):
		if n not in hash_map:
			res.append(n)
	return res

# time - O(n) space - O(1)
def findDisappearedNumbers(nums: list[int]) -> list[int]:
	res = []

	for n in nums:
		i = abs(n) - 1
		nums[i] = -abs(nums[i])

	for i in range(len(nums)):
		if nums[i] > 0:
			res.append(i + 1)
	return res