func subsets(nums []int) [][]int {
	result := [][]int{}

	var bt func([]int, int)

	bt = func(acc []int, start int) {
		if start > n {
			return
		}
		tmp := make([]int, len(acc))
		copy(tmp, acc)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			acc = append(acc, nums[i])
			bt(acc, i+1)
			acc = acc[:len(acc)-1]
		}
	}
	bt([]int{}, 0)
	return result
}
