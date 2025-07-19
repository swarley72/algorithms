func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	seen := NewSet[Coords]()
	// В очереди нужно хранить еще и номер раунда раскрутки изначально 1
	queue := makeDeque[QueueItem]([]QueueItem{})
	// пушим 0 в seen и очередь для "параллельного bfs"
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if mat[y][x] == 0 {
				coords := Coords{y, x}
				seen.Add(coords)
				queue.Append(QueueItem{coords, 1})
			}
		}
	}

	isValid := func(y, x int) bool {
		return x >= 0 && y >= 0 && x < n && y < m && !seen.Has(Coords{y, x})
	}

	distances := []Coords{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for queue.Size > 0 {
		levelSize := queue.Size
		for i := 0; i < levelSize; i++ {
			item := queue.PopLeft()
			coord := item.coords

			for _, d := range distances {
				newX := d.x + coord.x
				newY := d.y + coord.y

				if isValid(newY, newX) {
					mat[newY][newX] = item.round
					newCoord := Coords{newY, newX}
					seen.Add(newCoord)
					queue.Append(QueueItem{newCoord, item.round + 1})
				}
			}
		}
	}

	return mat
}

type Coords struct {
	y, x int
}

type QueueItem struct {
	coords Coords
	round  int
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
