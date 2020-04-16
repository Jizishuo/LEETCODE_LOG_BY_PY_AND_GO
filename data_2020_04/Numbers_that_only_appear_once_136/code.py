class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        # nums.sort()
        # n = list(set(nums[::2]) - set(nums[1::2]))[0]
        # return n

        return 2 * sum(set(nums)) - sum(nums)
