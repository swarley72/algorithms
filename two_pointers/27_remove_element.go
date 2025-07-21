func removeElement(nums []int, val int) int {
	p1 := 0

	for p2 := 0; p2 < len(nums); p2++ {
		if nums[p2] != val {
			nums[p1] = nums[p2]
			p1++
		}
	}
	return p1
}