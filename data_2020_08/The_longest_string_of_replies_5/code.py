class Solution:
    def longestPalindrome(self, s: str) -> str:
    #     if len(s)<2 or s==s[::-1]:
    #         return s
    #     max_len = 1
    #     num = s[0]
    #     for l in range(len(s)-1):
    #         if max_len>len(s[l+1:]):
    #             break
    #         for r in range(l+max_len, len(s)):
    #             if self.stt(s[l:r+1]) and r-l+1>max_len:
    #                 max_len = r-l+1
    #                 num = s[l:r+1]
    #
    #     return num
    #
    # def stt(self, s):
    #     return s == s[::-1]
        size = len(s)
        if size < 2:
            return s

        dp = [[False for _ in range(size)] for _ in range(size)]

        max_len = 1
        start = 0

        for i in range(size):
            dp[i][i] = True

        for j in range(1, size):
            for i in range(0, j):
                if s[i] == s[j]:
                    if j - i < 3:
                        dp[i][j] = True
                    else:
                        dp[i][j] = dp[i + 1][j - 1]
                else:
                    dp[i][j] = False

                if dp[i][j]:
                    cur_len = j - i + 1
                    if cur_len > max_len:
                        max_len = cur_len
                        start = i
        return s[start:start + max_len]


asss = "babab"
print(asss[1][2])
s = Solution().longestPalindrome(asss)
print(s)