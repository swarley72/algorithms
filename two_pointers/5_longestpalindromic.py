def longestPalindrome(s: str) -> str:
    l = 0
    r = len(s) - 1

    return s[0: r + 1]

print(longestPalindrome("babad"))

arr = [1, 2, 3, 4]

print(arr[1:3])