
func makeDeque[T any](items []T) *Deque[T] {
	head := &QueueNode[T]{}
	tail := &QueueNode[T]{}

	current := head
	for _, item := range items {
		node := &QueueNode[T]{Value: item}
		current.Next = node
		node.Prev = current
		current = node
	}
	current.Next = tail
	tail.Prev = current
	return &Deque[T]{Head: head, Tail: tail, Size: len(items)}
}

type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
	Prev  *QueueNode[T]
}

type Deque[T any] struct {
	Head *QueueNode[T]
	Tail *QueueNode[T]
	Size int
}

func (d *Deque[T]) Append(item T) {
	node := &QueueNode[T]{Value: item}
	node.Next = d.Tail
	node.Prev = d.Tail.Prev
	d.Tail.Prev.Next = node
	d.Tail.Prev = node
	d.Size++
}

func (d *Deque[T]) AppendLeft(item T) {
	node := &QueueNode[T]{Value: item}
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
