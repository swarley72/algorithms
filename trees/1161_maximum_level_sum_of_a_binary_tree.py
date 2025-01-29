from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def max_level_sum(root: TreeNode) -> int:
    result = 1
    max_sum = root.val
    queue = deque([(root, 1)])
    while queue:
        level_size = len(queue)
        tmp_sum = 0

        for _ in range(level_size):
            node, lvl = queue.popleft()
            tmp_sum += node.val

            if node.left:
                queue.append((node.left, lvl + 1))

            if node.right:
                queue.append((node.right, lvl + 1))
        if tmp_sum > max_sum:
            result = lvl
        max_sum = max(max_sum, tmp_sum)

    return result
