func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := nums[0]
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > last+1 {
			if start == last {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, last))
			}
			start = nums[i]
		}
		last = nums[i]
	}
	if start == last {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, last))
	}
	return result
}
