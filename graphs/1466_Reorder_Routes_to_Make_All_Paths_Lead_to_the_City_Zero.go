func minReorder(n int, connections [][]int) int {
	// 1 - создаем граф
	graph := map[int][]int{}
	// 2 - Тут храним старые направления (нужно для подсчета переворотов)
	oldDirections := NewSet[Direction]()
	// 3 - делаем граф двунаправленным чтобы ходить из нуля
	for _, con := range connections {
		a := con[0]
		b := con[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		// 4 - добавляем оригинальное направление
		oldDirections.Add(Direction{a, b})
	}
	// 5 - обычный обход графа dfs
	seen := NewSet[int]()
	seen.Add(0)
	reorders := 0
	stack := NewStack[int]()
	stack.Append(0)
	for len(stack) > 0 {
		city := stack.Pop()

		for _, ngh := range graph[city] {
			if !seen.Has(ngh) {
				seen.Add(ngh)
				stack.Append(ngh)

				// 6 - если оригинальное направление от нуля, но это переворот
				if oldDirections.Has(Direction{city, ngh}) {
					reorders++
				}
			}
		}
	}

	return reorders
}

// Так храним старые направления
type Direction struct {
	a, b int
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

