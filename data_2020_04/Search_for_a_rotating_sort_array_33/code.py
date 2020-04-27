class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return -1

        l= 0
        r = len(nums)-1

        while l<r:
            mid = (r-l)//2+l
            if nums[mid] == target:
                return mid
            elif nums[mid]>nums[l]:
                if nums[l] <= target <= nums[mid]:
                    r = mid
                else:
                    l = mid + 1
            else:
                if nums[mid+1] <= target <= nums[r]:
                    l = mid+1
                else:
                    r = mid
        return l if nums[l] == target else -1


