class Solution:
    def romanToInt(self, s: str) -> int:
        res = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        }
        ii = 0
        for index in range(len(s)-1):
            if res[s[index]] < res[s[index+1]]:
                ii -= res[s[index]]
            else:
                ii += res[s[index]]
        return ii + res[s[-1]]
