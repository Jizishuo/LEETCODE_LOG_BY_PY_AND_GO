class Solution:
    def longestPalindrome(self, s: str) -> int:
        d = {}
        for i in s:
            if i not in d.keys():
                d[i] = 1
            else:
                d[i]+=1

        res = 0
        # odd = 0
        # for i in d:
        #     if d[i] % 2 == 0:
        #         res +=d[i]
        #     elif d[i]%2==1:
        #         odd=1
        #         res +=d[i]-1
        # return res+1 if odd else res

        for i in d:
            res += (d[i]%2)*2

        if res == len(s):
            return res
        else:return res+1


            # import collections
            # # 1.统计各字符次数，eg:"ddsad":[3, 1, 1]
            # count = collections.Counter(s).values()
            # # 2.统计两两配对的字符总个数，eg: {"ddass":4,"ddsss":4}
            # x = sum([item // 2 * 2 for item in count if (item // 2 > 0)])
            # # 3.判断是否有没配对的单字符，有结果加一。 eg: {"ddss":4, "ddhjSS":4+1}-->{"ddss":4, "ddhjSS":5}
            # return x if x == len(s) else x + 1


