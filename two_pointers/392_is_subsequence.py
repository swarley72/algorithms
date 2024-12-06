def isSubsequence(s: str, t: str) -> bool:
        p = 0
        for i in range(len(s)):
            # if p == len(t):
            #     return True
            if s[i] == t[p]:
                p += 1

        return p == len(t)

