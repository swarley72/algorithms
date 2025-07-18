func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		prev := result[len(result)-1]
		start := intervals[i][0]
		end := intervals[i][1]

		if prev[1] >= start {
			prev[1] = max(prev[1], end)
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}