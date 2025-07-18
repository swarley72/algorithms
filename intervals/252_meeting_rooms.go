import "sort"

func meetingRooms(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if prev[1] > current[0] {
			return false
		}
		prev = current
	}
	return true
}