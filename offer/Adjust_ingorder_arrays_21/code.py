class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        # res = []
        # for i in nums:
        #     if i %2 == 0:
        #         res = [i]+res
        #     else:
        #         res = res + [i]
        # return res
        ll = len(nums)
        if ll == 0 or ll == 1:
            return nums
        l, r = 0, ll-1
        while l<r:
            if nums[l]%2==1 and nums[r]%2==1:
                l+=1
            elif nums[l]%2==0 and nums[r]%2==0:
                r-=1
            elif nums[l] % 2 == 1 and nums[r] % 2 == 0:
                l += 1
                r -=1
            elif nums[l] % 2 == 0 and nums[r] % 2 == 1:
                nums[l], nums[r] = nums[r], nums[l]
                l +=1
                r -= 1
        return nums
