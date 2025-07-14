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

type Pair struct {
	Left  *TreeNode
	Right *TreeNode
}

// ИТЕРАТИВНО (также через очередь)
func isSymmetric(root *TreeNode) bool {
	stack := Stack[Pair]{Pair{root.Left, root.Right}}

	for len(stack) > 0 {
		pair := stack.Pop()

		if pair.Left == nil && pair.Right == nil {
			continue
		}
		if pair.Left == nil || pair.Right == nil {
			return false
		}

		if pair.Left.Val != pair.Right.Val {
			return false
		}

		stack.Append(Pair{pair.Left.Right, pair.Right.Left})
		stack.Append(Pair{pair.Left.Left, pair.Right.Right})
	}

	return true
}

// РЕКУРСИВНО
func dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil || right == nil) || (left.Val != right.Val) {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}