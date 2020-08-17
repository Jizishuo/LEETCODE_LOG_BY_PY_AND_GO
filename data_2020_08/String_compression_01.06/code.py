class Solution:
    def compressString(self, S: str) -> str:
        cs = ' '
        i = 0
        for j in S:
            if j != cs[-1]:
                cs += str(i) + j
                i = 1
            else:
                i += 1
        return cs[2:] + str(i) if len(cs) - 2 < len(S) - 1 and i != 0 else S
