def longestPalindrome(s: str) -> str:
    currLen = 0
    res = ""
    for i in range(len(s)):
        # odd -> "bab"
        l,r = i,i
        while l >= 0 and r < len(s) and s[l] == s[r]:
            size = r - l + 1
            if size > currLen:
                currLen = size
                res = s[l:r+1]
            l -= 1
            r += 1

        # even -> "baab"
        l,r = i,i+1
        while l >= 0 and r < len(s) and s[l] == s[r]:
            size = r - l + 1
            if size > currLen:
                currLen = size
                res = s[l:r+1]
            l -= 1
            r += 1

    return res


            
            