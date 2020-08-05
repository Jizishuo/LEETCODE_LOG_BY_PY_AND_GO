class Solution:
    def rob(self, nums: List[int]) -> int:
        def ff(nums):
            c, p = 0, 0
            for i in nums:
                p, c = c, max(p+i, c)
            return c
        return max(ff(nums[:-1]), ff(nums[1:])) if len(nums) !=1 else nums[0]