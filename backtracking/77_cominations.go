func combine(n int, k int) [][]int {
	result := [][]int{}
	bt([]int{}, 1, n, k, &result)
	return result
}

func bt(acc []int, start, n, k int, result *[][]int) {
	if len(acc) == k {
		tmp := make([]int, k)
		copy(tmp, acc)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < n+1; i++ {
		acc = append(acc, i)
		bt(acc, i+1, n, k, result)
		acc = acc[0 : len(acc)-1]
	}
}