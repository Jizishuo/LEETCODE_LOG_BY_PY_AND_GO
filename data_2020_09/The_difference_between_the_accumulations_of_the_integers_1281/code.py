class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        a, b = 0, 1
        while n>0:
            d = n%10
            n//=10
            a +=d
            b*=d
        return b-a