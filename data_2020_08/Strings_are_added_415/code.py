class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        res = ""

        i, j, c = len(num1)-1, len(num2)-1, 0
        while i>=0 or j >=0:
            n1 = int(num1[i]) if i >=0 else 0
            n2 = int(num2[j]) if j >=0 else 0
            t = n1 + n2 + c
            c = t //10
            res = str(t%10)+res
            i, j = i-1, j-1
        return "1" + res if c else res