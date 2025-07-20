package graphs

func nearestExit(maze [][]byte, entrance []int) int {
	n := len(maze)
	m := len(maze[0])
	rowEnter := entrance[0]
	colEnter := entrance[1]

	seen := NewSet[Coord]()
	seen.Add(Coord{rowEnter, colEnter})

	queue := makeDeque[Cell]([]Cell{{rowEnter, colEnter, 0}})

	isValid := func(row, col int) bool {
		return row >= 0 && row < n && col >= 0 && col < m && maze[row][col] != '+' && !seen.Has(Coord{row, col})
	}
	isExit := func(row, col int) bool {
		return (row == 0 || row == n-1 || col == 0 || col == m-1) && (row != rowEnter || col != colEnter)
	}

	distances := []Coord{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			cell := queue.PopLeft()

			if isExit(cell.Row, cell.Col) {
				return cell.Steps
			}

			for _, d := range distances {
				newRow := cell.Row + d.Row
				newCol := cell.Col + d.Col

				if isValid(newRow, newCol) {
					seen.Add(Coord{newRow, newCol})
					queue.Append(Cell{newRow, newCol, cell.Steps + 1})
				}
			}

		}
	}

	return -1

}

type Coord struct {
	Row, Col int
}

type Cell struct {
	Row, Col, Steps int
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
