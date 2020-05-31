# class Solution:
#     def rotate(self, nums: List[int], k: int) -> None:
#         """
#         Do not return anything, modify nums in-place instead.
#         """
#         pass

k = 3
nums = [1,2,3,4,5,6,7,4]


# for _ in range(k):
#     nums = [nums[-1]] + nums[:-1]
# print(nums)
# return nums

n = len(nums)

k = k%len(nums)
print(k)
nums[:-k], nums[-k:] = nums[-k:], nums[:-k]
print(nums)

n = len(nums)
k %= n
nums[:] = nums[-k:] + nums[:-k]