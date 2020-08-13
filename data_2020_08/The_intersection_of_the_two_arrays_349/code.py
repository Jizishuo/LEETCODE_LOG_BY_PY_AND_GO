class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = []
        for i in set(nums1):
            for j in set(nums2):
                if i==j:
                    res.append(i)
        return res