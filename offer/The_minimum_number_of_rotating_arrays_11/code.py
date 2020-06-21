class Solution:
    def minArray(self, numbers: List[int]) -> int:

        l = 0
        r = len(numbers)-1
        while l< r:
            m = (l+r)//2
            m_v = numbers[m]
            if m_v>numbers[r]:
                l = m+1

            elif m_v<numbers[r]:
                r = m
            else: r -= 1
        return numbers[l]