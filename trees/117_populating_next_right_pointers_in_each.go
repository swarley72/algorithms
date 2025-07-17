func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := makeDeque[*Node]([]*Node{root})

	for queue.Size > 0 {
		levelSize := queue.Size
		var prev *Node

		for i := 0; i < levelSize; i++ {
			node := queue.PopLeft()
			if prev != nil {
				prev.Next = node
			}
			prev = node

			if node.Left != nil {
				queue.Append(node.Left)
			}

			if node.Right != nil {
				queue.Append(node.Right)
			}
		}
	}

	return root
}

// TYPES
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func makeDeque[T any](items []T) *Deque[T] {
	head := &Node2[T]{}
	tail := &Node2[T]{}

	current := head
	for _, item := range items {
		node := &Node2[T]{Value: item}
		current.Next = node
		node.Prev = current
		current = node
	}
	current.Next = tail
	tail.Prev = current
	return &Deque[T]{Head: head, Tail: tail, Size: len(items)}
}

type Node2[T any] struct {
	Value T
	Next  *Node2[T]
	Prev  *Node2[T]
}

type Deque[T any] struct {
	Head *Node2[T]
	Tail *Node2[T]
	Size int
}

func (d *Deque[T]) Append(item T) {
	node := &Node2[T]{Value: item}
	node.Next = d.Tail
	node.Prev = d.Tail.Prev
	d.Tail.Prev.Next = node
	d.Tail.Prev = node
	d.Size++
}

func (d *Deque[T]) AppendLeft(item T) {
	node := &Node2[T]{Value: item}
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
