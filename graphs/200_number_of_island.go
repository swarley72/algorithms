// Так хранятся координаты суши
type Tuple struct {
	y, x int
}

func numIslands(grid [][]byte) int {
	result := 0
	m := len(grid)
	n := len(grid[0])

	isValid := func(y, x int) bool {
		return (y >= 0 &&
			y < m &&
			x >= 0 &&
			x < n &&
			grid[y][x] == '1')
	}

	directions := []Tuple{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			result++
			stack := NewStack[Tuple]()
			stack.Append(Tuple{i, j})
			for len(stack) > 0 {
				tuple := stack.Pop()

				for _, dir := range directions {
					dx := dir.x + tuple.x
					dy := dir.y + tuple.y
					if isValid(dy, dx) {
						grid[dy][dx] = '0'
						stack.Append(Tuple{dy, dx})
					}
				}
			}
		}
	}

	return result
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

