class Solution:
    def searchRange(self, nums, target):
#         l = []
#         for i, v in enumerate(nums):
#             if v == target:
#                 print(i, v)
#                 l.append(i)
#         if len(l) == 0:
#             return [-1,-1]
#         else:
#             return [l[0], l[0]+len(l)-1]
#
# ds = Solution().searchRange([5,7,7,8,8,10], 8)
# print(ds)
        index = -1
        n = len(nums)
        l = 0
        r = n-1
        while l <=r:
            mid = (l+r) //2
            if nums[mid] == target:
                index = mid
                break
            elif nums[mid] <target:
                l = mid + 1
            else:
                r = mid - 1
        if index == -1:
            return [-1,-1]
        else:
            index1 = index
            while index1-1>-1 and nums[index-1] == target:
                index1 -= 1
            while index+1 <n and nums[index+1] == target:
                index += 1
            return [index1, index]