func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	seen := map[*Node]*Node{}
	seen[node] = &Node{Val: node.Val}

	stack := NewStack[*Node]()
	stack.Append(node)

	for len(stack) > 0 {
		v := stack.Pop()

		for _, ngh := range v.Neighbors {
			_, ok := seen[ngh]
			if !ok {
				seen[ngh] = &Node{Val: ngh.Val}
				stack.Append(ngh)
			}
			seen[v].Neighbors = append(seen[v].Neighbors, seen[ngh])
		}
	}

	return seen[node]
}

// TYPES
type Node struct {
	Val       int
	Neighbors []*Node
}

type Stack[T any] []T

// удаляет и возвращает последний элемент стека
func (s *Stack[T]) Pop() T {
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

// возвращает последний элемент стека
func (s Stack[T]) Top() T {
	return s[len(s)-1]
}

// добавляет элемент в стэк
func (s *Stack[T]) Append(val T) {
	*s = append(*s, val)
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}


