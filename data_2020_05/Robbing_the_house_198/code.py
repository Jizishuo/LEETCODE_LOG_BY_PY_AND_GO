class Solution:
    def rob(self, nums: List[int]) -> int:
        # n = len(nums)+1
        # db = [0]*(n-1)
        # db[0]=0
        # db[1] = nums[0]
        # for k in range(2, n+1):
        #     db[k] = max(db[k-1], nums[k-1]+db[k-2])
        #
        # return dp[n]

        p, c = 0, 0
        for i in nums:
            p, c = c, max(c, p+i)

        return c