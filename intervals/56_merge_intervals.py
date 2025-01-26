def merge(intervals: list[list[int]]) -> list[list[int]]:
    intervals.sort()

    result = [intervals[0]]

    for i in range(1, len(intervals)):
        a1, b1 = result[-1]
        a2, b2 = intervals[i]

        if max(a1, a2) <= min(b1, b2):
            result[-1] = [min(a1, a2), max(b1, b2)]
        else:
            result.append([a2, b2])

    return result
