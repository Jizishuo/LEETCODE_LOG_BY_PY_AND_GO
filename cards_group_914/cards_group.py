# -*- coding: utf-8 -*-
import math

class Solution:
    def hasGroupsSizeX(self, deck) -> bool:
        D = {}
        for x in deck:
            D[x] = D.setdefault(x, 0) + 1
        a = []
        for v in D.values():
            a.append(v)

        print(D, a)
        gcd = a[0]
        for i in range(1, len(a)):
            gcd = math.gcd(gcd, a[i])
        if gcd == 1:
            return False
        return True

"""
输入：[1,2,3,4,4,3,2,1]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
"""