class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        def func(n):
            return 0 if n<=0 else 1+ func(n & (n-1))
        return func(x^y)