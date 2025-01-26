from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def diameter_of_binary_tree(root: Optional[TreeNode]) -> int:
    d = 0

    def dfs(node):
        nonlocal d
        if not node:
            return 0

        left = dfs(node.left)
        right = dfs(node.right)

        d = max(d, left + right)

        return max(left, right) + 1

    dfs(root)
    return d
