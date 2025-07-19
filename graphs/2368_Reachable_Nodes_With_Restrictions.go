func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := map[int][]int{}

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := 0
	seen := NewSet[int]()
	for _, n := range restricted {
		seen.Add(n)
	}
	seen.Add(0)
	stack := NewStack[int]()
	stack.Append(0)

	for len(stack) > 0 {
		v := stack.Pop()
		visited++

		for _, ngh := range graph[v] {
			if !seen.Has(ngh) {
				seen.Add(ngh)
				stack.Append(ngh)
			}
		}

	}

	return visited

}

// TYPES
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
