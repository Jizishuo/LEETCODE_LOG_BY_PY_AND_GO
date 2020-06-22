class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n <0:
             x = 1/x
             n = -n
        res = 1
        while n:
            if n %2:
                res *=x
            x*=x
            n//=2
        return res