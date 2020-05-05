class Solution:
    def isHappy(self, n: int) -> bool:
        if (n==1):
            return True
        seen = {}
        def sumin(n:int):
            res = 0
            while n!=0:
                res += (n%10)**2
                n = n//10
            return res

        while n not in seen:
            seen[n]=True
            n=sumin(n)
            if n == 1:
                return True
        return False
