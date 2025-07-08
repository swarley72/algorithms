import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	l := 0
	r := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		leftNum := nums[l]
		rightNum := nums[r]
		if math.Abs(float64(leftNum)) > math.Abs(float64(rightNum)) {
			res[i] = leftNum * leftNum
			l++
		} else {
			res[i] = rightNum * rightNum
			r--
		}
	}

	return res
}
