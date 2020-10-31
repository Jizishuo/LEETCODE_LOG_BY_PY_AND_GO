class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        res = 0
        for i in range(len(nums)-1):
            for n in nums[i+1:]:
                if nums[i] == n:
                    res +=1
        return res