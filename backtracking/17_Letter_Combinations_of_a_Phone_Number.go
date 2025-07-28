package backtracking

func letterCombinations(digits string) []string {
	symMap := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	result := []string{}

	if len(digits) == 0 {
		return result
	}

	var bt func(word string, pos int)

	bt = func(word string, pos int) {
		if len(word) == len(digits) {
			result = append(result, word)
			return
		}

		for _, letter := range symMap[digits[pos]] {
			bt(word+letter, pos+1)
		}
	}

	bt("", 0)

	return result
}
