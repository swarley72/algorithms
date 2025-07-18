
const (
	L = "L"
	R = "R"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	direction := map[string]string{
		L: R,
		R: L,
	}
	currentDirection := L

	queue := makeDeque[*TreeNode]([]*TreeNode{root})

	for queue.Size > 0 {
		levelSize := queue.Size
		tmp := []int{}
		for i := 0; i < levelSize; i++ {
			var node *TreeNode
			if currentDirection == L {
				node = queue.PopLeft()
			} else {
				node = queue.Pop()
			}

			tmp = append(tmp, node.Val)

			if currentDirection == L {
				if node.Left != nil {
					queue.Append(node.Left)
				}
				if node.Right != nil {
					queue.Append(node.Right)
				}
			} else {
				if node.Right != nil {
					queue.AppendLeft(node.Right)
				}
				if node.Left != nil {
					queue.AppendLeft(node.Left)
				}
			}

		}
		currentDirection = direction[currentDirection]
		result = append(result, tmp)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
