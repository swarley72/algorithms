class Solution:
    def searchMatrix(self, matrix: list[list[int]], target: int) -> bool:
        def get_row() -> int:
            lo, hi = 0, len(matrix) - 1
            while lo <= hi:
                mid = lo + (hi - lo) // 2
                row = matrix[mid]
                if target >= row[0] and target <= row[-1]:
                    return mid
                elif target > row[-1]:
                    lo = mid + 1
                else:
                    hi = mid - 1
                
            return -1
        
        row_idx = get_row()
        if row_idx == -1:
            return False
        row = matrix[row_idx]
        lo, hi = 0, len(row) - 1
        while lo <= hi:
            mid = lo + (hi - lo) // 2

            if target == row[mid]:
                return True
            elif target > row[mid]:
                lo = mid + 1
            else:
                hi = mid - 1
        return False