import "unicode"

func isAlphanumeric(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	l := 0
	r := len(s) - 1

	for l < r {
		left := unicode.ToLower(rune(runes[l]))
		if !isAlphanumeric(unicode.ToLower(left)) {
			l++
			continue
		}

		right := unicode.ToLower(rune(runes[r]))
		if !isAlphanumeric(right) {
			r--
			continue
		}

		if left != right {
			return false
		}

		l++
		r--
	}
	return true
}
