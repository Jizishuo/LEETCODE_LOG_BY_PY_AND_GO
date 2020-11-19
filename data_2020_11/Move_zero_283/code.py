class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        idx_0 = 0
        for idx, num in enumerate(nums):
            if num == 0:
                continue
            if idx_0 == idx:
                idx_0 += 1
                continue
            nums[idx_0] = num
            nums[idx] = 0
            idx_0 +=1