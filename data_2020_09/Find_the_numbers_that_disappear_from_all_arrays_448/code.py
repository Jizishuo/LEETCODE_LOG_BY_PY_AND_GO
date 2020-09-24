class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        result = []
        for i in range(len(nums)):
            new_i = abs(nums[i])-1
            if nums[new_i]>0:
                nums[new_i] *= -1
        for j in range(len(nums)):
            if nums[j] >0:
                result.append(j+1)
        return result