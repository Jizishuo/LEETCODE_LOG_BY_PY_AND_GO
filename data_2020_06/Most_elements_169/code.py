class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        # l2 = len(nums)/2
        # aa = list(set(nums))
        # dd = {}
        # for i in aa:
        #     dd[i]=0
        #
        # for i in nums:
        #     if i in dd:
        #         dd[i] +=1
        #     if dd[i] > l2:
        #         return i
        m = 0
        c = 0
        for n in nums:
            if c == 0 :
                m = n
            if n == m:
                c += 1
            else:
                c -= 1
        return m