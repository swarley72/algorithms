func generateParenthesis(n int) []string {
	result := []string{}
	var bt func(acc string, open_cnt, close_cnt int)

	bt = func(acc string, open_cnt, close_cnt int) {
		if len(acc) == 2*n {
			result = append(result, acc)
			return
		}

		if open_cnt < n {
			bt(acc+"(", open_cnt+1, close_cnt)
		}
		if close_cnt < open_cnt {
			bt(acc+")", open_cnt, close_cnt+1)
		}

	}
	bt("", 0, 0)

	return result
}