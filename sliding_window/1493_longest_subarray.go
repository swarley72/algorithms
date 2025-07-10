// нужно вернуть длину подмассива состоящего из подряд идущих единиц минус один
func longestSubarray(nums []int) int {
	begin := 0
	windowState := 0
	result := 0

	for end := range len(nums) {
		// если находим ноль увеличиваем счетчик нулей
		if nums[end] == 0 {
			windowState++
		}

		// пока нулей больше одного идем до первого нуля
		for windowState > 1 {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}

		result = max(result, end-begin+1)
	}

	return result - 1
}

// [0, 0, 0, 0] - begin будет двигаться за end на каждом шаге в result будет 1, но по условию если нет едини нужно вернуть 0 --> 0
// [1, 0, 0] - на первой итерации result == 1 (0 - 0 + 1), на второй будет тоже 1 потому что end = 2, begin = 2 --> 0
// [1, 1, 0]