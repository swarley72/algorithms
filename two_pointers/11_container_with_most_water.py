def max_area(height: list[int]) -> int:
    l = 0
    r = len(height) - 1
    result = 0

    while l < r:
        width = r - l
        left_height = height[l]
        right_height = height[r]
        min_height = min(left_height, right_height)
        result = max(result, min_height * width)

        if left_height < right_height:
            l += 1
        else:
            r -= 1

    return result
