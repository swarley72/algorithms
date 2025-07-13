type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

func isValid(s string) bool {
	pMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := Stack[rune]{}
	for _, c := range s {
		_, ok := pMap[c]
		if ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack.Pop()
			if c != pMap[last] {
				return false
			}
		}

	}

	return len(stack) == 0
}
