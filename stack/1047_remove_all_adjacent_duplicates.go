import "strings"

type Stack[T any] []T

// удаляет и возвращает последний элемент стека
func (s *Stack[T]) Pop() T {
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

// возвращает последний элемент стека
func (s Stack[T]) Top() T {
	return s[len(s)-1]
}

// добавляет элемент в стэк
func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func removeDuplicates(s string) string {
	stack := Stack[string]{}

	for _, char := range s {
		if len(stack) != 0 && stack.Top() == string(char) {
			stack.Pop()
		} else {
			stack.Push(string(char))
		}
	}

	return strings.Join(stack, "")
}
