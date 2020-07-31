class Solution:
    def findMagicIndex(self, nums: List[int]) -> int:
        # if len(nums) <= 1:
        #     if nums[0]==0:
        #         return 0
        #     else:return -1
        # if nums[0]>nums[-1]:
        #     nums = nums[::-1]
        #
        # for i in range(len(nums)):
        #     if nums[i]==i:
        #         return i
        # return -1
        if not nums:return -1
        if nums[0]==0:return 0
        l, r = 0, len(nums)
        while l<r:
            if nums[l]>l:
                l = nums[l]
            elif nums[l]==l:
                return l
            else:
                l+=1
        return -1