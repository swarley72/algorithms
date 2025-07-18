type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(node *TreeNode, d *int) int {
	if node == nil {
		return 0
	}

	left := maxDepth(node.Left, d)
	right := maxDepth(node.Right, d)
	*d = max(*d, left+right)

	return 1 + max(left, right)
}
func diameterOfBinaryTree(root *TreeNode) int {
	d := 0

	maxDepth(root, &d)
	return d
}