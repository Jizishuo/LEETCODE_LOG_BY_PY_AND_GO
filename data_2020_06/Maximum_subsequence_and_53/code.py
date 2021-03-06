class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # temp = nums[0]
        # maxx = temp
        # n = len(nums)
        # for i in range(1, n):
        #     if maxx + nums[i] > nums[i]:
        #         maxx = max(maxx, temp+nums[i])
        #         temp = temp+nums[i]
        #     else:
        #         maxx = max(maxx, temp, tmp + nums[i], nums[i])
        #         temp = nums[i]
        # return maxx
        tmp = nums[0]
        max_ = tmp
        n = len(nums)
        for i in range(1, n):
            # 当当前序列加上此时的元素的值大于tmp的值，说明最大序列和可能出现在后续序列中，记录此时的最大值
            if tmp + nums[i] > nums[i]:
                max_ = max(max_, tmp + nums[i])
                tmp = tmp + nums[i]
            else:
                # 当tmp(当前和)小于下一个元素时，当前最长序列到此为止。以该元素为起点继续找最大子序列,
                # 并记录此时的最大值
                max_ = max(max_, tmp, tmp + nums[i], nums[i])
                tmp = nums[i]
        return max_




