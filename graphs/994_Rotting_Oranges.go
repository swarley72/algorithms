package graphs

func orangesRotting(grid [][]int) int {
	cntTomato := 0
	timer := 0
	n := len(grid)
	m := len(grid[0])

	queue := makeDeque[Tomato]([]Tomato{})
	rotten := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 0 {
				cntTomato++
			}
			if grid[i][j] == 2 {
				rotten++
				queue.Append(Tomato{i, j, 0})
			}
		}
	}

	isValid := func(row, col int) bool {
		return row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == 1
	}

	directions := []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			tomato := queue.PopLeft()
			timer = max(timer, tomato.Time)

			for _, d := range directions {
				newRow := tomato.Row + d.Row
				newCol := tomato.Col + d.Col
				if isValid(newRow, newCol) {
					rotten++
					grid[newRow][newCol] = 2
					queue.Append(Tomato{newRow, newCol, tomato.Time + 1})
				}

			}
		}
	}

	if cntTomato == rotten {
		return timer
	}
	return -1
}

type Coord struct {
	Row, Col int
}

type Tomato struct {
	Row, Col, Time int
}

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
