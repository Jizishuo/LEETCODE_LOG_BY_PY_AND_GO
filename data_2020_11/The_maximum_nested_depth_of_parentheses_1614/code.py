class Solution:
    def maxDepth(self, s: str) -> int:
        max_x = 0
        t = 0
        for i in s:
            if i == '(':
                t+=1

            if max_x <=t:
                max_x = t
            if i == ')':
                t-=1
        return max_x