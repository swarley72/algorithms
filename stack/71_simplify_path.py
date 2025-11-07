class Solution:
    def simplifyPath(self, path: str) -> str:
        path = path.split("/")

        stack = []

        for p in path:
            if not p or p == ".":
                continue
            if p == "..":
                if stack:
                    stack.pop()
            else:
                stack.append(p)

        return "/" + "/".join(stack)
