class Solution:
    def reverse(self, x: int) -> int:
        y = int(str(x)[::-1]) if x >= 0 else -int(str(x)[:0:-1])
        return y if -2 ** 31 < y < 2 ** 31 - 1 else 0