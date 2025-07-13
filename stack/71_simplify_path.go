import (
	"strings"
	"fmt"
)

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

func simplifyPath(path string) string {
	splitted := strings.Split(path, "/")
	stack := Stack[string]{}

	for _, p := range splitted {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(stack) > 0 {
				stack.Pop()
			}
		} else {
			stack.Push(p)
		}
	}

	return "/" + strings.Join(stack, "/")
}