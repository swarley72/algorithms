def combine( n: int, k: int) -> list[list[int]]:
    result = []

    def backtrack(acc, start):
        if len(acc) == k:
            result.append(acc[:])
            return

        for num in range(start, n + 1):
            if num not in acc:
                acc.append(num)
                backtrack(acc, num + 1)
                acc.pop()

    backtrack([], 1)
    return result
