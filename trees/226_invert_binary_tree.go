type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// рекурсивно
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// итеративно
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := Stack[*TreeNode]{root}

	for len(stack) > 0 {
		node := stack.Pop()
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			stack.Append(node.Left)
		}

		if node.Right != nil {
			stack.Append(node.Right)
		}
	}

	return root
}
