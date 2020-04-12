class Solution:
    def reverseWords(self, s: str) -> str:
        ss = s.split(" ")[::-1]
        print(ss)
        return " ".join(ss)


dd = Solution().reverseWords("the sky is blue")
print(dd)