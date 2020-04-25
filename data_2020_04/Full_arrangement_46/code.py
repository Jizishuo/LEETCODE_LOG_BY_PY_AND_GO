class Solution:
    def permute(self, nums):
        if len(nums) <= 1:
            return [nums]

        res = []
        for i, num in enumerate(nums):
            rest = nums[:i] + nums[i + 1:]
            for rest_num in self.permute(rest):
                res.append([num] + rest_num)
        print(res)
        return res




Solution().permute([1,2,3])