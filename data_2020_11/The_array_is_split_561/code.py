class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        nums.sorts()
        res = 0
        for i in range(0, len(nums)-1, 2):
            res += nums[i]
        return res