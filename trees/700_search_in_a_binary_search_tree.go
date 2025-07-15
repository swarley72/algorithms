
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

// ИТЕРАТИВНО
func searchBST(root *TreeNode, val int) *TreeNode {
	stack := Stack[*TreeNode]{root}

	for len(stack) > 0 {
		node := stack.Pop()

		if node == nil {
			continue
		}

		if node.Val == val {
			return node
		}

		if node.Val > val {
			stack.Append(node.Left)
		} else {
			stack.Append(node.Right)
		}

	}

	return nil
}

// РЕКУРСИВНО
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}