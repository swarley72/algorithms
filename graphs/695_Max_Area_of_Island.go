func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	isValid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m && grid[y][x] == 1
	}
	directions := []Coords{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	maxArea := 0

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 1 {
				tmpArea := 1
				grid[y][x] = 0
				stack := NewStack[Coords]()
				stack.Append(Coords{x, y})
				for len(stack) > 0 {
					coords := stack.Pop()
					for _, d := range directions {
						newX := coords.x + d.x
						newY := coords.y + d.y
						if isValid(newX, newY) {
							grid[newY][newX] = 0
							stack.Append(Coords{newX, newY})
							tmpArea++
						}
					}
				}
				maxArea = max(maxArea, tmpArea)
			}
		}
	}

	return maxArea
}

type Coords struct {
	x, y int
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

