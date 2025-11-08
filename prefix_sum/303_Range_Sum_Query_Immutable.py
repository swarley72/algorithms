class NumArray:
    def __init__(self, nums: list[int]):
        self.prefix_sum = [0]
        acc = 0
        for n in nums:
            acc += n
            self.prefix_sum.append(acc)

    def sumRange(self, left: int, right: int) -> int:
        return self.prefix_sum[right + 1] - self.prefix_sum[left]


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
