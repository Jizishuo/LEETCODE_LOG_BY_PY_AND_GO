class Solution:
    def fib(self, n: int) -> int:
        if n ==1:
            return 0
        if n==2:
            return 1
        if n ==3:
            return 2

        return self.fib(n-1)+self.fib(n-2)