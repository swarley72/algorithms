func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegree := make([]int, n)

	for i := 0; i < len(edges); i++ {
		to := edges[i][1]
		inDegree[to]++
	}

	result := []int{}
	for idx, n := range inDegree {
		if n == 0 {
			result = append(result, idx)
		}
	}
	return result
}