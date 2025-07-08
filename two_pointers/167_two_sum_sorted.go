func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		twoSum := numbers[l] + numbers[r]
		if twoSum == target {
			return []int{l + 1, r + 1}
		}

		if twoSum < target {
			l++
		} else {
			r--
		}

	}
	return []int{}
}
