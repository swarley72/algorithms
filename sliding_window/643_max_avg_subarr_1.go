import "math"

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	windowState := 0.0
	result := math.Inf(-1) // float64

	for end := range len(nums) {
		windowState += float64(nums[end])
		windowSize := end - begin + 1
		if windowSize == k {
			result = max(result, windowState)
			windowState -= float64(nums[begin])
			begin++
		}
	}

	return result / float64(k)
}