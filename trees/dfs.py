def dfs(node):
    if not node:
        return
    # print(node.val) # pre order
    dfs(node.left)
    # print(node.val) # in order
    dfs(node.right)
    # print(node.val) # post order
    return