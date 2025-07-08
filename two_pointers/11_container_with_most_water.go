func maxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1

	for l < r {
		width := r - l
		leftH := height[l]
		rightH := height[r]

		res = max(res, width*min(leftH, rightH))

		if leftH < rightH {
			l++
		} else {
			r--
		}

	}
	return res
}
