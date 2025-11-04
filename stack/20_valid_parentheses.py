class Solution:
    def isValid(self, s: str) -> bool:
        p_map = {"}": "{", ")": "(", "]": "["}
        stack = []

        for p in s:
            if p in p_map:
                if not stack:
                    return False
                if stack.pop() != p_map[p]:
                    return False
            else:
                stack.append(p)

        if stack:
            return False
        return True
