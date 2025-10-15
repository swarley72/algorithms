def minNumber(nums1: list[int], nums2: list[int]) -> int:
	set1 = set(nums1)
	set2 = set(nums2)

	if inter := set1 & set2:
		return min(inter)
	
	return min(int(f"{min(set1)}{min(set2)}"), int(f"{min(set2)}{min(set1)}"))