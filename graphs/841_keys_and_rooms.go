func canVisitAllRooms(rooms [][]int) bool {
	seen := NewSet[int]()
	seen.Add(0)
	stack := NewStack[int]()
	stack.Append(0)

	for len(stack) > 0 {
		room := stack.Pop()

		for _, key := range rooms[room] {
			if !seen.Has(key) {
				seen.Add(key)
				stack.Append(key)
			}
		}
	}

	return len(seen.elements) == len(rooms)

}

// TYPES
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

