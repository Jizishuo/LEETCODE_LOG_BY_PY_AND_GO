class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        d = {}
        for i in nums1:
            if i in d:
                d[i]+=1
            else:
                d[i] =1
        res = []
        for i in nums2:
            if i in d and d[n]>0:
                res.append(n)
                d[i] -=1
        return res

