class Solution:
    def waysToStep(self, n: int) -> int:
        a,b,c = 4,2,1
        if n <3:
            return n
        if n == 3:
            return 4
        for i in range(n-3):
            a,b,c = (a+b+c)%1000000007,a,b
        return a