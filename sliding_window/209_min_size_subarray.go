import "math"

func minSubArrayLen(target int, nums []int) int {
	begin := 0
	result := math.Inf(1) // float64
	windowState := 0

	for end := range len(nums) {
		windowState += nums[end]

		for windowState >= target {
			windowSize := end - begin + 1
			result = min(result, float64(windowSize)) // приведение типа
			windowState -= nums[begin]
			begin++
		}
	}

	if result == math.Inf(1) {
		return 0
	}

	return int(result) // приведение типа
}