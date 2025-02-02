from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_symmetric_stack(root: TreeNode) -> bool:
    stack = [(root.left, root.right)]
    while stack:
        left, right = stack.pop()
        if not left and not right:
            continue
        if not left or not right or left.val != right.val:
            return False

        stack.append((left.left, right.right))
        stack.append((left.right, right.left))

    return True


def is_symmetric_queue(root: TreeNode) -> bool:
    queue = deque([(root.left, root.right)])
    while queue:
        left, right = queue.popleft()
        if not left and not right:
            continue
        if not left or not right or left.val != right.val:
            return False

        queue.append((left.left, right.right))
        queue.append((left.right, right.left))

    return True

def is_symmetric_recursive(self, root: TreeNode) -> bool:
    def dfs(left, right):
        if not left and not right:
            return True

        if not left or not right or left.val != right.val:
            return False

        return dfs(left.left, right.right) and dfs(left.right, right.left)

    return dfs(root.left, root.right)
