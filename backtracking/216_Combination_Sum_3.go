func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	var bt func(acc []int, start, path_sum int)

	bt = func(acc []int, start, path_sum int) {
		if len(acc) == k {
			if path_sum == n {
				tmp := make([]int, k)
				copy(tmp, acc)
				result = append(result, tmp)
			}
			return
		}
		if path_sum >= n {
			return
		}

		for i := start; i <= 9; i++ {
			acc = append(acc, i)
			bt(acc, i+1, path_sum+i)
			acc = acc[0 : len(acc)-1]
		}

	}
	bt([]int{}, 1, 0)

	return result
}
