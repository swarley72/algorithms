class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:

        def get_curr_index(_str, i):
            skip = 0
            while i >= 0:
                if _str[i] == "#":
                    skip += 1
                elif skip > 0:
                   skip -= 1
                else:
                    break
                i -= 1 

            return i
            

        p1 = len(s) - 1
        p2 = len(t) - 1

        while p1 >= 0 or p2 >=0:
            p1 = get_curr_index(s, p1)
            p2 = get_curr_index(t, p2)

            if p1 < 0 and p2 < 0:
                return True
            if p1 < 0 or p2 < 0:
                return False
            if s[p1] != t[p2]:
                return False
            
            p1 -= 1
            p2 -= 1

        return True
        