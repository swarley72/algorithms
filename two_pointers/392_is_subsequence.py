def isSubsequence(s: str, t: str) -> bool:
        p1 = 0

        for c in t:
            # если уже вышли за границу
            if p1 == len(s):
                return True
            if c == s[p1]:
                p1 += 1
        
        return p1 == len(s)
