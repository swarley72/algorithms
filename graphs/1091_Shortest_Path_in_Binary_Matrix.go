func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	if grid[0][0] != 0 {
		return -1
	}
	grid[0][0] = 1

	queue := makeDeque[Item]([]Item{Item{Coord{0, 0}, 1}})

	isValid := func(row, col int) bool {
		return row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == 0
	}

	directions := []Coord{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			v := queue.PopLeft()

			if v.Coord.Row == n-1 && v.Coord.Col == m-1 {
				return v.Steps
			}

			for _, d := range directions {
				newRow := v.Coord.Row + d.Row
				newCol := v.Coord.Col + d.Col
				if isValid(newRow, newCol) {
					grid[newRow][newCol] = 1

					queue.Append(Item{Coord{newRow, newCol}, v.Steps + 1})
				}
			}
		}

	}

	return -1
}

type Coord struct {
	Row, Col int
}

type Item struct {
	Coord
	Steps int
}

// TYPES
func makeDeque[T any](items []T) *Deque[T] {
	head := &Node[T]{}
	tail := &Node[T]{}

	current := head
	for _, item := range items {
		node := &Node[T]{Value: item}
		current.Next = node
		node.Prev = current
		current = node
	}
	current.Next = tail
	tail.Prev = current
	return &Deque[T]{Head: head, Tail: tail, Size: len(items)}
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type Deque[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (d *Deque[T]) Append(item T) {
	node := &Node[T]{Value: item}
	node.Next = d.Tail
	node.Prev = d.Tail.Prev
	d.Tail.Prev.Next = node
	d.Tail.Prev = node
	d.Size++
}

func (d *Deque[T]) AppendLeft(item T) {
	node := &Node[T]{Value: item}
	node.Next = d.Head.Next
	d.Head.Next.Prev = node
	d.Head.Next = node
	d.Size++
}

func (d *Deque[T]) Pop() T {
	popValue := d.Tail.Prev.Value
	node := d.Tail.Prev
	node.Prev.Next = d.Tail
	d.Tail.Prev = node.Prev
	d.Size--
	return popValue
}

func (d *Deque[T]) PopLeft() T {
	popValue := d.Head.Next.Value

	node := d.Head.Next
	node.Next.Prev = d.Head
	d.Head.Next = node.Next
	d.Size--
	return popValue
}
