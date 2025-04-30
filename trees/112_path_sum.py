from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def has_path_sum(root: Optional[TreeNode], target_sum: int) -> bool:
    if not root:
        return False

    if not root.left and not root.right:
        return root.val == target_sum

    left = has_path_sum(root.left, target_sum - root.val)
    right = has_path_sum(root.right, target_sum - root.val)

    return left or right
