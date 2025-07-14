type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// РЕКУРСИВНО
func dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil || right == nil) || (left.Val != right.Val) {
		return false
	}

	return dfs(left.Left, right.Left) && dfs(left.Right, right.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p, q)
}