
import "math"

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
	node *TreeNode
	min  float64
	max  float64
}

func isValidBST(root *TreeNode) bool {
	stack := Stack[Item]{Item{root, math.Inf(-1), math.Inf(1)}}

	for len(stack) > 0 {
		item := stack.Pop()
		node := item.node
		min := item.min
		max := item.max

		if node == nil {
			continue
		}

		if float64(node.Val) >= max || float64(node.Val) <= min {
			return false
		}

		if node.Left != nil {
			stack.Append(Item{node.Left, min, float64(node.Val)})
		}
		if node.Right != nil {
			stack.Append(Item{node.Right, float64(node.Val), max})
		}
	}

	return true
}