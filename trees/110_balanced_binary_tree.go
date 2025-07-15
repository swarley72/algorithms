import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(height(node.Left), height(node.Right))
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.Left)
	right := height(root.Right)

	if math.Abs(float64(left)-float64(right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}