type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

type Item struct {
	node   *TreeNode
	target int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := Stack[Item]{Item{root, 0}}
	for len(stack) > 0 {
		item := stack.Pop()
		node := item.node
		currentSum := item.target + node.Val

		if node.Left == nil && node.Right == nil && currentSum == targetSum {
			return true
		}

		if node.Left != nil {
			stack.Append(Item{node.Left, currentSum})

		}

		if node.Right != nil {
			stack.Append(Item{node.Right, currentSum})
		}

	}

	return false
}