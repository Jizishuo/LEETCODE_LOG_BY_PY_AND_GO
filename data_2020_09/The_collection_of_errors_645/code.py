class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        a = sum(nums) - sum(set(nums))
        c = 0
        for i, num in enumerate(nums, 1):
            c ^=(i^num)
        b = c^a
        return [a,b]