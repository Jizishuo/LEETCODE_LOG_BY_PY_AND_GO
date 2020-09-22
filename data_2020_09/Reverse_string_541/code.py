class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        res, flag = '', True
        for i in range(0, len(s), k):
            res += s[i:i+k][::-1] if flag else s[i:i+k]
            flag = not flag
        return res