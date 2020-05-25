class Solution:
    def validPalindrome(self, s: str) -> bool:

        def search(chars, times):
            begin, end = 0, len(chars)-1
            while begin<end:
                if chars[begin] != chars[end]:
                    if times == 0:
                        return False
                    return search(chars[begin:end], times - 1) or search(chars[begin + 1:end + 1], times - 1)
                begin+=1
                end-=1
            return True
        return search(s, 1)