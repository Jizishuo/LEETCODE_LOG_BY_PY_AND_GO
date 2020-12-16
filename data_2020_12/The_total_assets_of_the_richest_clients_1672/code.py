class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        max_n = 0
        for i in accounts:
            a = sum(i)
            if a > max_n:
                max_n = a
        return max_n