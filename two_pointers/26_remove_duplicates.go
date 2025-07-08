func removeDuplicates(nums []int) int {
	l := 0

	for r := range len(nums) {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}