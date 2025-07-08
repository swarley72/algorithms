func isSubsequence(s string, t string) bool {
	p1 := 0

	for p2 := range len(t) {
		// если уже вышли за границу
		if p1 == len(s) {
			return true
		}
		if s[p1] == t[p2] {
			p1++
		}
	}

	return p1 == len(s)
}