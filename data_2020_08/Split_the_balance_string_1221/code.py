class Solution:
    def balancedStringSplit(self, s: str) -> int:
        if s[0] == 'L':
            tmp=[1]
        else:tmp = [-1]

        for i in range(1, len(s)):
            if s[i]=='L':
                tmp.append(tmp[-1]+1)
            else:
                tmp.append(tmp[-1]-1)
        return tmp.count(0)
