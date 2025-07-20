import "math"

type COLOR int

const (
	RED COLOR = iota
	BLUE
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	// создаем мапу мап
	graph := map[COLOR]map[int][]int{
		RED:  make(map[int][]int),
		BLUE: make(map[int][]int),
	}
	nextColor := map[COLOR]COLOR{
		RED:  BLUE,
		BLUE: RED,
	}

	// заполняем мапы по цветам
	for _, edge := range redEdges {
		a := edge[0]
		b := edge[1]
		graph[RED][a] = append(graph[RED][a], b)
	}

	for _, edge := range blueEdges {
		a := edge[0]
		b := edge[1]
		graph[BLUE][a] = append(graph[BLUE][a], b)
	}
	// заполняем массив бесконечностями чтобы сравнивать по min
	result := make([]float64, n)
	for i := range result {
		result[i] = math.Inf(1)
	}
	// в нулевую ноду всегда можем дойти
	result[0] = 0

	// для нулевой вершины добавляем с обоими цветами
	queue := makeDeque[Item]([]Item{Item{0, RED, 0}, Item{0, BLUE, 0}})

	seen := NewSet[V]()
	seen.Add(V{0, RED})
	seen.Add(V{0, BLUE})

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			node := queue.PopLeft()
			// меняем результат так как мы сюда уже дошли
			result[node.Val] = min(result[node.Val], float64(node.Path))

			// берем соседей с текущего цвета
			for _, ngh := range graph[node.Color][node.Val] {
				v := V{ngh, nextColor[node.Color]}
				// проверяем что мы еще не были в вершине другого цвета, т.к. нам нужно менять цвета ребер
				if !seen.Has(v) {
					seen.Add(v)
					queue.Append(Item{ngh, nextColor[node.Color], node.Path + 1})
				}
			}

		}
	}

	// просто меняем типы на int
	newResult := make([]int, n, n)
	for idx, n := range result {
		if n == math.Inf(1) {
			newResult[idx] = -1
		} else {
			newResult[idx] = int(n)
		}
	}
	return newResult
}

type V struct {
	Val   int
	Color COLOR
}

type Item struct {
	Val   int
	Color COLOR
	Path  int
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
