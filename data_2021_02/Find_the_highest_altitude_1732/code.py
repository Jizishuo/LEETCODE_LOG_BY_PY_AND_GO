class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        mt = 0
        a = 0
        for i in gain:
            a += i
            if a > mt:
                mt = a
        return mt
