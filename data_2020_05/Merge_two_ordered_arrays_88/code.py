class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        if m == 0:
            for i in range(n):
                nums1[i] = nums2[i]

        for x in nums2:
            for i in range(m):
                if x < nums1[i]:
                    nums1.insert(i, x)
                    nums1.pop()
                    m = m + 1
                    break
                if x >=nums1[m-1]:
                    nums1.insert(m,x)
                    nums1.pop()
                    m = m+1
                    break