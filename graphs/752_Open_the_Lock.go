import "strings"

func getNextComb(comb string) []string {
	combs := make([]string, 0, 8)
	runes := []rune(comb)
	for idx, r := range runes {
		num := int(r - '0')

		nextNum := (num + 1) % 10
		nextRunes := make([]rune, 4)
		copy(nextRunes, runes)
		nextRunes[idx] = rune(nextNum + '0')
		combs = append(combs, getCodeString(nextRunes))

		prevNum := (num + 9) % 10
		prevRunes := make([]rune, 4)
		copy(prevRunes, runes)
		prevRunes[idx] = rune(prevNum + '0')

		combs = append(combs, getCodeString(prevRunes))
	}

	return combs
}

func getCodeString(nums []rune) string {
	var b strings.Builder
	for _, r := range nums {
		b.WriteString(string(r))
	}
	return b.String()
}

func openLock(deadends []string, target string) int {
	seen := NewSet[string]()
	for _, deadend := range deadends {
		seen.Add(deadend)
	}
	if seen.Has("0000") {
		return -1
	}

	queue := makeDeque[Comb]([]Comb{{"0000", 0}})

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			comb := queue.PopLeft()
			if comb.Val == target {
				return comb.Steps
			}

			for _, newComb := range getNextComb(comb.Val) {
				if !seen.Has(newComb) {
					seen.Add(newComb)
					queue.Append(Comb{newComb, comb.Steps + 1})
				}
			}
		}
	}

	return -1
}

type Comb struct {
	Val   string
	Steps int
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
