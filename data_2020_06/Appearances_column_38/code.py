
from itertools import groupby
class Solution:
    def countAndSay(self, n: int) -> str:
        # result = '1'
        # for i in range(1, n):
        #     result = ''.join([str(len(list(g))) + k for k, g in groupby(result)])
        # return result

        if n == 1:
            return '1'
        elif n == 2:
            return '11'
        else:
            count = 1
            res = ''
            s = self.countAndSay(n - 1)
            for i in range(1, len(s)):
                if s[i] == s[i - 1]:
                    count += 1
                else:
                    res = res + str(count) + s[i - 1]
                    count = 1
            res = res + str(count) + s[i]
            return res



