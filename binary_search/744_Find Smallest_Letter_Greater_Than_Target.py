class Solution:
    def nextGreatestLetter(self, letters: list[str], target: str) -> str:
        l, r = 0, len(letters) - 1
        res = ""
        while l <= r:
            mid = l + (r -l) // 2
            if letters[mid] > target:
                res = letters[mid]
                r = mid - 1
            else:
                l = mid + 1
        if not res:
            return letters[0]
        return res