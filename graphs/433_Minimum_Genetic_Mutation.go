import "strings"

func getNextGenes(gen string) []string {
	bank := []string{}
	runes := []rune(gen)
	genes := []rune{'A', 'C', 'G', 'T'}

	for idx, r := range runes {
		for _, g := range genes {
			if r != g {
				nextGen := make([]rune, 8)
				copy(nextGen, runes)
				nextGen[idx] = g
				bank = append(bank, getGeneString(nextGen))
			}
		}
	}

	return bank
}

func getGeneString(gens []rune) string {
	var b strings.Builder
	for _, r := range gens {
		b.WriteString(string(r))
	}
	return b.String()
}

func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := NewSet[string]()
	for _, g := range bank {
		bankSet.Add(g)
	}

	seen := NewSet[string]()
	seen.Add(startGene)

	queue := makeDeque[Gene]([]Gene{Gene{startGene, 0}})

	for queue.Size > 0 {
		levelSize := queue.Size

		for i := 0; i < levelSize; i++ {
			gene := queue.PopLeft()
			if gene.Val == endGene {
				return gene.Steps
			}
			for _, newGene := range getNextGenes(gene.Val) {
				if bankSet.Has(newGene) && !seen.Has(newGene) {
					seen.Add(newGene)
					queue.Append(Gene{newGene, gene.Steps + 1})
				}
			}

		}
	}

	return -1
}

type Gene struct {
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
