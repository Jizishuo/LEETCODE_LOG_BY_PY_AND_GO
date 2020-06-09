class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        # 不管这个数在不在里面，直接append
        nums.append(target)
        # 然后再排序
        nums.sort()
        # 最后返回查找的index
        return nums.index(target)

