def sorted_squares(self, nums: list[int]) -> list[int]:
        n = len(nums)
        res = [0] * n
        
        l = 0
        r = n - 1

        for i in range(n -1, -1, -1):
            left_num = nums[l]
            right_num = nums[r]
            if abs(left_num) > abs(right_num):
                res[i] = left_num*left_num
                l += 1
            else:
                 res[i] = right_num*right_num
                 r -= 1
        return res
