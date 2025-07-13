type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Рекурсивно
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

// итеративно со стеком
type Item struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max_depth := 1

	stack := Stack[Item]{Item{root, 1}}

	for len(stack) > 0 {
		item := stack.Pop()
		max_depth = max(max_depth, item.Depth)

		if item.Node.Right != nil {
			stack.Append(Item{item.Node.Right, item.Depth + 1})
		}
		if item.Node.Left != nil {
			stack.Append(Item{item.Node.Left, item.Depth + 1})
		}
	}

	return max_depth
}

// Типы для стэка
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