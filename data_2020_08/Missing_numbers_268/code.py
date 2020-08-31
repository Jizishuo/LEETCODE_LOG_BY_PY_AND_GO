class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        miss = len(nums)
        for i, num in enumerate(nums):
            miss ^=i^num
        return miss