class Solution:
    def judgeSquareSum(self, c: int) -> bool:

        l, r = 0, int(c**0.5)
        while l<=r:
            ll = l*l
            rr = r*r
            if ll + rr ==c:
                return True
            elif ll + rr > c:
                r -=1
            elif ll + rr < c:
                l+=1
        return False