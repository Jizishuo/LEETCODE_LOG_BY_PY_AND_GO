class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.lower()

        l = []
        for i in s:
            if i.isalnum():
                l.append(i)
        return l==l[::-1]