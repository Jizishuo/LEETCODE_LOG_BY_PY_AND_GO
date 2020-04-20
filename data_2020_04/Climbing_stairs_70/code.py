class Solution:
    def climbStairs(self, n: int) -> int:
        one = 1
        two = 2
        res = 0
        for i in range(2, n):
            res = one+two

            one = two
            two = res
        return max(n, res)