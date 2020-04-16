class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        nums.sort(key=lambda x: x == val)
        return len(nums) - nums.count(val)

"""
    # i为不同元素的数组的长度
    i = 0
    for j in range(0, len(nums)):
        # 如果nums[j]不等于val, 则将nums[j]赋值给nums[i]即可, i自增
        if nums[j] != val:
            nums[i] = nums[j]
            i += 1
    return i

"""