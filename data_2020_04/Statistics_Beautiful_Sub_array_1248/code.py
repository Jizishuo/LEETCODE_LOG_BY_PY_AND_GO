
class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        n = len(nums)
        odd = [-1]
        ans = 0
        for i in range(n):
            if nums[i] % 2 == 1:
                odd.append(i)
        odd.append(n)
        print(odd)
        for i in range(1, len(odd) - k):
            ans += (odd[i] - odd[i - 1]) * (odd[i + k] - odd[i + k - 1])
        return ans


#
# nums = [1,1,2,1,1]
# k = 3

nums = [2,2,2,1,2,2,1,2,2,2]
k = 2
Solution().numberOfSubarrays(nums, k)
