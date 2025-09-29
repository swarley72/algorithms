def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
	p1 = m - 1
	p2 = n - 1

	p3 = len(nums1) - 1

	while p1 >= 0 and p2 >= 0:
		num_1 = nums1[p1]
		num_2 = nums2[p2]
		if num_1 > num_2:
			nums1[p3] = num_1
			p1 -= 1
		else:
			nums1[p3] = num_2
			p2 -= 1
		p3 -= 1
	
	while p2 >= 0:
		nums1[p3] = nums2[p2]
		p2 -= 1
		p3 -= 1
	
