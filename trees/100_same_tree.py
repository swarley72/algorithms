from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# ---------------------------------
def is_same_tree_recursive(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    if not p and not q:
        return True
    if not p or not q:
        return False

    if p.val != q.val:
        return False

    left = is_same_tree_recursive(p.left, q.left)
    right = is_same_tree_recursive(p.right, q.right)

    return left and right


def is_same_tree_stack(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    stack = [(p, q)]
    while stack:
        l_node, r_node = stack.pop()

        if not l_node and not r_node:
            continue
        if not l_node or not r_node:
            return False

        if l_node.val != r_node.val:
            return False
        stack.append((l_node.left, r_node.left))
        stack.append((l_node.right, r_node.right))

    return True