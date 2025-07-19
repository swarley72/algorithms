func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	graph := map[int][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				graph[i] = append(graph[i], j)
			}
		}
	}

	result := 0
	seen := NewSet[int]()
	for i := 0; i < n; i++ {
		if !seen.Has(i) {
			result++
			seen.Add(i)
			stack := NewStack[int]()
			stack.Append(i)
			for len(stack) > 0 {
				c := stack.Pop()
				for _, ngh := range graph[c] {
					if !seen.Has(ngh) {
						seen.Add(ngh)
						stack.Append(ngh)
					}
				}
			}
		}

	}

	return result

}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

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
func (s *Stack[T]) Append(val T) {
	*s = append(*s, val)
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}
