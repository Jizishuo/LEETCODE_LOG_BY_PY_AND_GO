class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        # s = set()
        # for i in nums:
        #     if i in s:
        #         return i
        #     else:
        #         s.add(i)

        l = len(nums)
        tmp = 0

        for i in range(l):
            while nums[i] != i:
                if nums[i] == nums[nums[i]]:
                    return num[i]
                tmp = nums[i]
                nums[i] = nums[tmp]
                nums[tmp] = tmp

