func longestOnes(nums []int, k int) int {
	begin := 0
	windowState := 0 // количество встреченых нулей
	result := 0

	for end := range len(nums) {
		if nums[end] == 0 {
			windowState++
		}

		// если был хотя бы один ноль то отсуюад выйдем
		for windowState > k {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}

		result = max(result, end-begin+1)
	}

	return result
}